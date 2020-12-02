package data

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// ReadData provides a simple way to read data from a file data source
func ReadData(s io.Reader) (r []int, err error) {
	r = []int{}

	scanner := bufio.NewScanner(s)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), ",")
		for _, v := range values {
			vi, _ := strconv.Atoi(v)
			r = append(r, vi)
		}
	}

	err = scanner.Err()

	return
}
