package domain

type ID string

type Rule interface {
	Validate() error
}

func CheckRule(rule Rule) error {
	return rule.Validate()
}
