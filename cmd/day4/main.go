package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/easterthebunny/advent2020/internal/data"
)

var (
	source = flag.String("f", "./cmd/day4/data.txt", "data file")
)

func main() {
	flag.Parse()

	file, err := os.Open(*source)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	d, _ := data.ReadPassportData(file)
	count := data.CountValidPassports(*d)

	fmt.Printf("valid passports in data file: %d\n", count)
}
