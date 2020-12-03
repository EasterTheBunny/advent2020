package data

import (
	"sort"

	"github.com/easterthebunny/advent2020/internal/types"
)

// CountValidPasswords counts all valid passwords in the list
func CountValidPasswords(list *[]PasswordDataEntry, tp types.PasswordRuleType) int {
	count := 0

	for _, e := range *list {
		rule, err := types.NewPasswordRule(e.Rule, tp)
		if err != nil {
			return 0
		}

		if rule.Validate(types.Password(e.Value)) {
			count++
		}
	}

	return count
}

// ComputeExpenseReport finds two numbers in the data source that add up to
// 2020 and returns the product of the two numbers.
func ComputeExpenseReport(data []int, sum, dimension int) (product int) {
	combos := combinations(dimension, len(data)-1)
	for _, c := range combos {
		d := make([]int, len(c), len(c))
		for x, i := range c {
			d[x] = data[i]
		}

		s := sumList(d...)
		if s == sum {
			product = productList(d...)
		}
	}

	return
}

func sumList(inputs ...int) (sum int) {
	for _, x := range inputs {
		sum = sum + x
	}

	return
}

func productList(inputs ...int) (product int) {
	product = 1
	for _, x := range inputs {
		product = product * x
	}

	return
}

type barrel struct {
	limit int
	start int
	value int
	next  *barrel
}

func (b *barrel) cycle() bool {
	if b.value < b.limit {
		b.value++
		return false
	}

	/*
		b.start++
		if b.start > b.limit {
			b.start = b.limit
		}
	*/

	b.value = b.start
	if b.next != nil {
		return b.next.cycle()
	}

	return true
}

func (b *barrel) report(values []int) []int {
	r := append(values, b.value)
	if b.next != nil {
		return b.next.report(r)
	}
	return r
}

func combinations(size, limit int) (combos [][]int) {
	result := &barrel{limit: limit, value: 0, start: 0, next: nil}
	for i := 0; i < size-1; i++ {
		result = &barrel{limit: limit, value: 0, start: 0, next: result}
	}

	rmap := map[string]bool{}
	for !result.cycle() {
		b := result.report([]int{})
		if unique(b) {
			sort.Slice(b, func(i, j int) bool {
				return b[i] > b[j]
			})
			key := ""
			for _, k := range b {
				key = key + string(k)
			}

			ok := rmap[key]
			if !ok {
				rmap[key] = true
				combos = append(combos, b)
			}

		}
	}

	return
}

func unique(d []int) bool {
	values := map[int]bool{}

	for _, v := range d {
		_, ok := values[v]
		if !ok {
			values[v] = true
		} else {
			return false
		}
	}

	return true
}
