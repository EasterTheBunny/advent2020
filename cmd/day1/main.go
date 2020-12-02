package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/easterthebunny/advent2020/internal/data"
)

var (
	source = flag.String("f", "./data.txt", "data file")
)

func main() {
	flag.Parse()

	file, err := os.Open(*source)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	d, err := data.ReadData(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("product of 2 entries: %d\n", data.ComputeExpenseReport(d, 2020, 2))
	fmt.Printf("product of 3 entries: %d\n", data.ComputeExpenseReport(d, 2020, 3))
}
