package jsonutil

type segmentKind int

const (
	segmentKey segmentKind = iota
	segmentIndex
)

type TokenPathSegment struct {
	kind  segmentKind
	key   string
	index int
}
