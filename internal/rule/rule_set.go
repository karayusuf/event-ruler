package rule

type RuleSet interface {
	Rules() []*Rule
}

type StandardRuleSet struct {
	rules []*Rule
}

func NewRuleSet(rules []*Rule) RuleSet {
	return &StandardRuleSet{
		rules: rules,
	}
}

func (rs *StandardRuleSet) Rules() []*Rule {
	return rs.rules
}
