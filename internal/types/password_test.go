package types

import "testing"

func TestCountRuleValidate(t *testing.T) {
	type tst struct {
		min, max int
		value    rune
		password Password
		expected bool
	}

	testData := []tst{
		{min: 1, max: 3, value: 'a', password: "abcde", expected: true},
		{min: 1, max: 3, value: 'b', password: "cdefg", expected: false},
		{min: 2, max: 9, value: 'c', password: "ccccccccc", expected: true}}

	for _, test := range testData {
		r := CountRule{min: test.min, max: test.max, value: test.value}
		v := r.Validate(test.password)

		if v != test.expected {
			t.Errorf("unexpected password validation result %v; expected %v", v, test.expected)
		}
	}
}

func TestPositionRuleValidate(t *testing.T) {
	type tst struct {
		min, max int
		value    rune
		password Password
		expected bool
	}

	testData := []tst{
		{min: 1, max: 3, value: 'a', password: "abcde", expected: true},
		{min: 1, max: 3, value: 'b', password: "cdefg", expected: false},
		{min: 2, max: 9, value: 'c', password: "ccccccccc", expected: false}}

	for i, test := range testData {
		r := PositionRule{pos1: test.min, pos2: test.max, value: test.value}
		v := r.Validate(test.password)

		if v != test.expected {
			t.Errorf("test index %d: unexpected password validation result %v; expected %v", i, v, test.expected)
		}
	}
}
