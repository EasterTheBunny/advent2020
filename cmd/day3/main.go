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
	source = flag.String("f", "./cmd/day3/data.txt", "data file")
	out    = flag.String("o", "./cmd/day3/data.copy.txt", "data file")
)

func main() {
	flag.Parse()

	file, err := os.Open(*source)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m, err := data.ReadMapData(file)
	if err != nil {
		log.Fatal(err)
	}

	type slptst struct {
		path  types.PathFunc
		trees int
	}

	slopes := []slptst{
		{
			path: func(p types.Position) types.Position {
				return types.Position{X: p.X + 1, Y: p.Y + 1}
			},
			trees: 0,
		},
		{
			path: func(p types.Position) types.Position {
				return types.Position{X: p.X + 3, Y: p.Y + 1}
			},
			trees: 0,
		},
		{
			path: func(p types.Position) types.Position {
				return types.Position{X: p.X + 5, Y: p.Y + 1}
			},
			trees: 0,
		},
		{
			path: func(p types.Position) types.Position {
				return types.Position{X: p.X + 7, Y: p.Y + 1}
			},
			trees: 0,
		},
		{
			path: func(p types.Position) types.Position {
				return types.Position{X: p.X + 1, Y: p.Y + 2}
			},
			trees: 0,
		},
	}

	product := 1
	start := types.Position{X: 0, Y: 0}
	b := &types.Toboggan{
		Position: start,
		Map:      m}

	for x, exp := range slopes {
		b.Position = start
		err = nil

		for ; err != types.ErrOffMap; err = b.Move(exp.path) {
			if err == types.ErrTreeError {
				slopes[x].trees++
			}
		}

		product = product * slopes[x].trees
		fmt.Printf("trees in path %d: %d;\n", x, slopes[x].trees)
	}

	fmt.Printf("product of all trees: %d\n", product)
}
