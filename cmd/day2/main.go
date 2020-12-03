package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/easterthebunny/advent2020/internal/data"
	"github.com/easterthebunny/advent2020/internal/types"
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

	fmt.Printf("valid passwords in data file for count rule: %d\n", data.CountValidPasswords(d, types.CountRuleType))
	fmt.Printf("valid passwords in data file for position rule: %d\n", data.CountValidPasswords(d, types.PositionRuleType))
}
