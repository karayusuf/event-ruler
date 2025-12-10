package jsonutil

type segmentKind int

const (
	segmentKey segmentKind = iota
	segmentIndex
)

type tokenPathSegment struct {
	kind  segmentKind
	key   string
	index int
}
