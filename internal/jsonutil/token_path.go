package jsonutil

import (
	"strconv"
	"strings"
)

type tokenPath struct {
	segments []tokenPathSegment
}

func (t *tokenPath) String() string {
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

func (t *tokenPath) pushKey(key string) {
	t.segments = append(t.segments, tokenPathSegment{kind: segmentKey, key: key})
}

func (t *tokenPath) pushIndex(index int) {
	t.segments = append(t.segments, tokenPathSegment{kind: segmentIndex, index: index})
}

func (t *tokenPath) pop() {
	length := len(t.segments)
	if length > 0 {
		t.segments = t.segments[:length-1]
	}
}
