package domain

type Rule interface {
	Validate() error
}

func CheckRule(rule Rule) error {
	return rule.Validate()
}
