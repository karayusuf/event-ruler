package rule

import "strings"

type Rule interface {
	GetPath() string
	GetPathSegments() []string
	Matches(value Value) bool
}

type StandardRule struct {
	path         string
	pathSegments []string
	matches      func(value Value) bool
}

func NewRule(pathSegments []string, matches func(value Value) bool) Rule {
	return &StandardRule{
		path:         strings.Join(pathSegments, "."),
		pathSegments: pathSegments,
		matches:      matches,
	}
}

func (r *StandardRule) GetPath() string {
	return r.path
}

func (r *StandardRule) GetPathSegments() []string {
	return r.pathSegments
}

func (r *StandardRule) Matches(value Value) bool {
	return r.matches(value)
}
