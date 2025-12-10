package jsonutil

import "fmt"

type SegmentKind int

const (
	SegmentKey SegmentKind = iota
	SegmentIndex
)

type PathSegment interface {
	Key() string // for object keys

	private() // seals the interface
}

type PathSegmentKey struct {
	key string
}

func NewPathSegmentKey(key string) PathSegment {
	return &PathSegmentKey{
		key: key,
	}
}

// Key implements PathSegment.
func (p *PathSegmentKey) Key() string {
	return p.key
}

// private implements PathSegment.
func (p *PathSegmentKey) private() {
}

type PathSegmentIndex struct {
	index int
}

func NewPathSegmentIndex(key string, index int) PathSegment {
	return &PathSegmentIndex{
		index: index,
	}
}

// Key implements PathSegment.
func (p *PathSegmentIndex) Key() string {
	return fmt.Sprintf("%d", p.index)
}

// private implements PathSegment.
func (p *PathSegmentIndex) private() {
}
