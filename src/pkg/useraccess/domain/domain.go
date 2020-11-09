package domain

import (
	"github.com/ke-pan/modular-monolith-with-ddd-in-go/src/pkg/useraccess/domain/userregistrastion/rule"
)

func CheckRule(rule rule.Rule) error {
	return rule.Validate()
}
