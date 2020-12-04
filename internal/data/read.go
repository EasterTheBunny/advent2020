package data

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"

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

// ReadMapData creates a new Map from a data source
func ReadMapData(s io.Reader) (*types.Map, error) {
	points := []types.Position{}

	scanner := bufio.NewScanner(s)
	line := 0
	width := 0
	for scanner.Scan() {
		trees, w := parseEncodedMapLine(line, scanner.Text())
		width = w
		points = append(points, trees...)
		line++
	}

	return types.NewMap(points, line-1, width), nil
}

func parseEncodedMapLine(y int, l string) ([]types.Position, int) {
	r := []types.Position{}
	line := []byte(l)
	width := 0

	for x, b := range line {
		if b == byte('#') {
			r = append(r, types.Position{X: x, Y: y})
		}
		width = x
	}

	return r, width + 1
}
