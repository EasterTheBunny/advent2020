package types

// Password is an abstraction that allows password rules to be applied
// to validate the password.
type Password string

// NewPasswordRule creates a password rule since all fields are unexported
func NewPasswordRule(min, max int, value rune) *PasswordRule {
	return &PasswordRule{min: min, max: max, value: value}
}

// PasswordRule represents a rule that applies to a password
type PasswordRule struct {
	min   int
	max   int
	value rune
}

// Validate checks whether or not a password is valid for this rule.
func (r *PasswordRule) Validate(p Password) bool {
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

// CheckRules runs validation on the password with all rules passed in.
func (p *Password) CheckRules(rules []PasswordRule) bool {
	for _, r := range rules {
		if !r.Validate(*p) {
			return false
		}
	}

	return true
}
