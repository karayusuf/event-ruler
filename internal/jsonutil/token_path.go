package jsonutil

import (
	"strconv"
	"strings"
)

type TokenPath struct {
	segments []TokenPathSegment
}

func (t *TokenPath) Clone() *TokenPath {
	if t == nil {
		return nil
	}

	// new slice with its own backing array
	segs := make([]TokenPathSegment, len(t.segments))
	copy(segs, t.segments)

	return &TokenPath{
		segments: segs,
	}
}

func (t *TokenPath) String() string {
	var path strings.Builder
	path.WriteByte('$')

	for _, segment := range t.segments {
		switch segment.kind {
		case segmentKey:
			path.WriteByte('.')
			path.WriteString(segment.key)
		case segmentIndex:
			path.WriteByte('[')
			path.WriteString(strconv.Itoa(segment.index))
			path.WriteByte(']')
		}
	}

	return path.String()
}

func (t *TokenPath) pushKey(key string) {
	t.segments = append(t.segments, TokenPathSegment{kind: segmentKey, key: key})
}

func (t *TokenPath) pushIndex(index int) {
	t.segments = append(t.segments, TokenPathSegment{kind: segmentIndex, index: index})
}

func (t *TokenPath) pop() {
	length := len(t.segments)
	if length > 0 {
		t.segments = t.segments[:length-1]
	}
}
