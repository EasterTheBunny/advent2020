package data

import (
	"strings"
	"testing"
)

func TestComputeExpenseReport(t *testing.T) {
	type tst struct {
		list      []int
		dimension int
		sum       int
		expected  int
	}

	testData := []tst{
		{list: []int{1721, 979, 366, 299, 675, 1456}, dimension: 2, sum: 2020, expected: 514579},
		{list: []int{1721, 979, 366, 299, 675, 1456}, dimension: 3, sum: 2020, expected: 241861950}}

	for _, test := range testData {
		result := ComputeExpenseReport(test.list, test.sum, test.dimension)

		if result != test.expected {
			t.Errorf("unexpected product %d; expected %d", result, test.expected)
		}
	}
}

func TestSumList(t *testing.T) {
	type tst struct {
		list     []int
		expected int
	}

	testData := []tst{
		{list: []int{1, 1, 1}, expected: 3},
		{list: []int{0, 0, 1}, expected: 1},
		{list: []int{1, 2, 3}, expected: 6}}

	for _, test := range testData {
		r := sumList(test.list...)
		if r != test.expected {
			t.Errorf("unexpected sum %d; expected %d", r, test.expected)
		}
	}
}

func TestProductList(t *testing.T) {
	type tst struct {
		list     []int
		expected int
	}

	testData := []tst{
		{list: []int{1, 1, 1}, expected: 1},
		{list: []int{0, 0, 1}, expected: 0},
		{list: []int{1, 2, 3}, expected: 6}}

	for _, test := range testData {
		r := productList(test.list...)
		if r != test.expected {
			t.Errorf("unexpected product %d; expected %d", r, test.expected)
		}
	}
}

func TestCombinations(t *testing.T) {
	type tst struct {
		size     int
		limit    int
		expected int
	}

	testData := []tst{
		{size: 2, limit: 3, expected: 6},
		{size: 2, limit: 4, expected: 10},
		{size: 2, limit: 9, expected: 45},
		{size: 3, limit: 3, expected: 4}}

	for _, test := range testData {
		c := combinations(test.size, test.limit)

		if len(c) != test.expected {
			t.Errorf("unexpected number of combinations found %d; expected %d", len(c), test.expected)
		}
	}
}

func TestUnique(t *testing.T) {
	type uniqueTest struct {
		values   []int
		expected bool
	}

	testData := []uniqueTest{
		{
			values:   []int{1, 1, 2, 3},
			expected: false},
		{
			values:   []int{1, 2, 3, 4},
			expected: true}}

	for _, test := range testData {
		result := unique(test.values)

		if result != test.expected {
			t.Errorf("uniqueness test failed for values %v; expected %v; got %v", test.values, test.expected, result)
		}
	}
}

func TestCountValidPassports(t *testing.T) {
	data := `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`

	r := strings.NewReader(data)
	p, _ := ReadPassportData(r)
	c := CountValidPassports(*p)

	if c != 2 {
		t.Errorf("unexpected passport count %d; expected %d", c, 2)
	}
}
