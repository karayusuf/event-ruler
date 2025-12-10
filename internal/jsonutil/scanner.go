package jsonutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type Scanner struct {
}

// CallbackOnLeaf is called for every scalar leaf value.
// pathSegments is a slice representing the path to the leaf.
//   - Object keys: "foo", "bar"
//   - Array indices: "0", "1", ...
//
// Example path: foo.bar[2].baz → []string{"foo", "bar", "2", "baz"}
type CallbackOnLeaf func(tokenPath *tokenPath, token json.Token) bool

var (
	// internal sentinel used for early termination when callback returns false
	errStop = errors.New("scan stopped by callback")
)

// Scan walks the JSON structure and calls onLeaf for every scalar leaf.
// It returns a slice of all leaf tokens encountered (in visit order).
//
// If onLeaf returns false at any point, scanning stops early.
func (s *Scanner) Scan(reader io.Reader, onLeaf CallbackOnLeaf) error {
	decoder := json.NewDecoder(reader)
	decoder.UseNumber()

	cursor := &cursor{
		currentPath:    &tokenPath{},
		decoder:        decoder,
		callbackOnLeaf: onLeaf,
	}

	if err := cursor.scanRoot(); err != nil {
		if errors.Is(err, errStop) || errors.Is(err, io.EOF) {
			return nil
		}
		return err
	}

	// Ensure there are no trailing tokens after a valid JSON value.
	if decoder.More() {
		return fmt.Errorf("unexpected extra data after top-level JSON value")
	}

	return nil
}

type cursor struct {
	currentPath    *tokenPath
	decoder        *json.Decoder
	callbackOnLeaf CallbackOnLeaf
}

func (c *cursor) scanRoot() error {
	// Read the first token (top-level value)
	tok, err := c.decoder.Token()
	if err != nil {
		return err
	}
	return c.scanValueWithToken(tok)
}

// scanValue reads and processes a value starting at the next token.
func (c *cursor) scanValue() error {
	tok, err := c.decoder.Token()
	if err != nil {
		return err
	}
	return c.scanValueWithToken(tok)
}

// scanValueWithToken processes a value when we've already consumed its first token.
func (c *cursor) scanValueWithToken(tok json.Token) error {
	switch v := tok.(type) {
	case json.Delim:
		switch v {
		case '{':
			return c.scanObject()
		case '[':
			return c.scanArray()
		default:
			return fmt.Errorf("unexpected delimiter %q", v)
		}
	default:
		// Scalar leaf
		if c.callbackOnLeaf != nil {
			if !c.callbackOnLeaf(c.currentPath, tok) {
				return errStop
			}
		}
		return nil
	}
}

func (c *cursor) scanObject() error {
	// {} or {"k": v, ...}
	for c.decoder.More() {
		// Read key
		tok, err := c.decoder.Token()
		if err != nil {
			return err
		}

		key, ok := tok.(string)
		if !ok {
			return fmt.Errorf("invalid token %T (%v): object key must be a string", tok, tok)
		}

		c.currentPath.pushKey(key)
		if err := c.scanValue(); err != nil {
			if errors.Is(err, errStop) {
				return err
			}
			return err
		}
		c.currentPath.pop()
	}

	// Consume closing '}'
	tok, err := c.decoder.Token()
	if err != nil {
		return err
	}
	d, ok := tok.(json.Delim)
	if !ok || d != '}' {
		return fmt.Errorf("expected '}', got %v", tok)
	}

	return nil
}

func (c *cursor) scanArray() error {
	// [] or [v1, v2, ...]
	index := 0

	for c.decoder.More() {
		c.currentPath.pushIndex(index)
		if err := c.scanValue(); err != nil {
			if errors.Is(err, errStop) {
				return err
			}
			return err
		}
		index++
		c.currentPath.pop()
	}

	// Consume closing ']'
	tok, err := c.decoder.Token()
	if err != nil {
		return err
	}
	d, ok := tok.(json.Delim)
	if !ok || d != ']' {
		return fmt.Errorf("expected ']', got %v", tok)
	}

	return nil
}
