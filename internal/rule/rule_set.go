package rule

import (
	"encoding/json"
	"strings"

	"github.com/karayusuf/event-ruler/internal/jsonutil"
)

type RuleSet struct {
	rules []Rule
}

func Parse(jsonRules string) (RuleSet, error) {
	ruleSet := RuleSet{}
	reader := strings.NewReader(jsonRules)

	err := jsonutil.Scan(reader, func(tokenPath *jsonutil.TokenPath, token json.Token) bool {
		// path := tokenPath
		return true
	})

	return ruleSet, err
}

func NewRuleSet(rules []Rule) RuleSet {
	return RuleSet{
		rules: rules,
	}
}

func (rs RuleSet) Rules() []Rule {
	return rs.rules
}
