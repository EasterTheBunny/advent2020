package types

import (
	"errors"
	"regexp"
	"strconv"
)

// Password is an abstraction that allows password rules to be applied
// to validate the password.
type Password string

// PasswordRuleType is a value that indicates the rule to be used
type PasswordRuleType int

const (
	// CountRuleType is the rule type that maps to a CountRule
	CountRuleType PasswordRuleType = iota
	// PositionRuleType is the rule type that maps to a PositionRule
	PositionRuleType
)

// NewPasswordRule creates a password rule since all fields are unexported
func NewPasswordRule(rule string, tp PasswordRuleType) (PasswordRule, error) {
	switch tp {
	case CountRuleType:
		r := &CountRule{raw: rule, Type: tp}
		ok := r.Parse(nil)
		if !ok {
			return nil, errors.New("rule parse error")
		}
		return r, nil
	case PositionRuleType:
		r := &PositionRule{raw: rule, Type: tp}
		ok := r.Parse(nil)
		if !ok {
			return nil, errors.New("rule parse error")
		}
		return r, nil
	default:
		return nil, errors.New("unknown password rule type")
	}
}

// PasswordRule represents a rule that applies to a password
type PasswordRule interface {
	Validate(Password) bool
	Parse(*string) bool
}

// CountRule represents a rule that applies to a password
type CountRule struct {
	Type  PasswordRuleType
	raw   string
	min   int
	max   int
	value rune
}

// Validate checks whether or not a password is valid for this rule.
func (r *CountRule) Validate(p Password) bool {
	counts := map[rune]int{}

	for _, r := range p {
		_, ok := counts[r]
		if !ok {
			counts[r] = 0
		}

		counts[r]++
	}

	count, ok := counts[r.value]
	if !ok && r.min > 0 {
		return false
	}

	if count < r.min || count > r.max {
		return false
	}

	return true
}

// Parse implements the PasswordRule interface for the count rule type
func (r *CountRule) Parse(rule *string) bool {
	if rule != nil {
		r.raw = *rule
	}
	regexp := regexp.MustCompile(`(?P<Min>\d+)-(?P<Max>\d+)\s(?P<Value>[a-zA-Z]{1})`)
	matches := regexp.FindStringSubmatch(r.raw)

	if len(matches) != 4 {
		return false
	}

	min, _ := strconv.Atoi(matches[1])
	max, _ := strconv.Atoi(matches[2])
	runeSlice := []rune(matches[3])

	r.min = min
	r.max = max
	r.value = runeSlice[0]

	return true
}

// PositionRule represents a rule that looks for a value at specific positions
type PositionRule struct {
	Type       PasswordRuleType
	raw        string
	pos1, pos2 int
	value      rune
}

// Validate checks whether or not a password is valid for this rule.
func (r *PositionRule) Validate(p Password) bool {
	if len(p) < r.pos2 {
		return false
	}

	if rune(p[r.pos1-1]) != r.value && rune(p[r.pos2-1]) != r.value {
		return false
	}

	if rune(p[r.pos1-1]) == r.value && rune(p[r.pos2-1]) == r.value {
		return false
	}

	return true
}

// Parse implements the PasswordRule interface for the position rule type
func (r *PositionRule) Parse(rule *string) bool {
	if rule != nil {
		r.raw = *rule
	}
	regexp := regexp.MustCompile(`(?P<Min>\d+)-(?P<Max>\d+)\s(?P<Value>[a-zA-Z]{1})`)
	matches := regexp.FindStringSubmatch(r.raw)

	if len(matches) != 4 {
		return false
	}

	min, _ := strconv.Atoi(matches[1])
	max, _ := strconv.Atoi(matches[2])
	runeSlice := []rune(matches[3])

	r.pos1 = min
	r.pos2 = max
	r.value = runeSlice[0]

	return true
}

// CheckRules runs validation on the password with all rules passed in.
func (p *Password) CheckRules(rules []PasswordRule) bool {
	for _, r := range rules {
		if !r.Validate(*p) {
			return false
		}
	}

	return true
}
