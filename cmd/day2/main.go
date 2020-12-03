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

	d, err := data.ReadPasswordData(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("valid passwords in data file: %d\n", data.CountValidPasswords(d))
}
