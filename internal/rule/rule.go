package rule

import "strings"

type RuleSet interface {
	GetRules() []*Rule
}

type Rule interface {
	GetPath() string
	GetPathSegments() []string
	Matches(value Value) bool
}

type SimpleRule struct {
	path         string
	pathSegments []string
	matches      func(value Value) bool
}

func NewSimpleRule(pathSegments []string, matches func(value Value) bool) Rule {
	return &SimpleRule{
		path:         strings.Join(pathSegments, "."),
		pathSegments: pathSegments,
		matches:      matches,
	}
}

func (r *SimpleRule) GetPath() string {
	return r.path
}

func (r *SimpleRule) GetPathSegments() []string {
	return r.pathSegments
}

func (r *SimpleRule) Matches(value Value) bool {
	return r.matches(value)
}
