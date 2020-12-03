package data

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
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
	Rule  string
	Value string
}

// ReadPasswordData creates a slice of password data entries
func ReadPasswordData(s io.Reader) (*[]PasswordDataEntry, error) {
	entries := []PasswordDataEntry{}

	scanner := bufio.NewScanner(s)
	for scanner.Scan() {
		opts := strings.Split(scanner.Text(), ": ")
		if len(opts) != 2 {
			return nil, errors.New("line parse error")
		}

		entries = append(entries, PasswordDataEntry{Rule: opts[0], Value: opts[1]})
	}

	err := scanner.Err()
	if err != nil {
		return nil, err
	}

	return &entries, nil
}
