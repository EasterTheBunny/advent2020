package data

import (
	"bufio"
	"io"
	"regexp"
	"strconv"

	"github.com/easterthebunny/advent2020/internal/types"
)

// ReadData provides a simple way to read data from a file data source
func ReadData(s io.Reader) (r []int, err error) {
	r = []int{}

	scanner := bufio.NewScanner(s)
	for scanner.Scan() {
		vi, _ := strconv.Atoi(scanner.Text())
		r = append(r, vi)
	}

	err = scanner.Err()

	return
}

// PasswordDataEntry is a line item in the password data file
type PasswordDataEntry struct {
	Rule  *types.PasswordRule
	Value types.Password
}

// ReadPasswordData creates a slice of password data entries
func ReadPasswordData(s io.Reader) (*[]PasswordDataEntry, error) {
	entries := []PasswordDataEntry{}

	scanner := bufio.NewScanner(s)
	for scanner.Scan() {
		r := regexp.MustCompile(`(?P<Min>\d+)-(?P<Max>\d+)\s(?P<Value>[a-zA-Z]{1}): (?P<Password>[a-zA-Z]+)`)
		matches := r.FindStringSubmatch(scanner.Text())

		min, _ := strconv.Atoi(matches[1])
		max, _ := strconv.Atoi(matches[2])
		runeSlice := []rune(matches[3])

		rule := types.NewPasswordRule(min, max, runeSlice[0])
		password := types.Password(matches[4])

		entries = append(entries, PasswordDataEntry{Rule: rule, Value: password})
	}

	err := scanner.Err()
	if err != nil {
		return nil, err
	}

	return &entries, nil
}
