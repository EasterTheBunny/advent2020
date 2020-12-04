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

	m, err := data.ReadMapData(file)
	if err != nil {
		log.Fatal(err)
	}

	b := &types.Toboggan{
		Position: types.Position{X: 0, Y: 0},
		Map:      m}

	moveFunc := func(p types.Position) types.Position {
		return types.Position{X: p.X + 3, Y: p.Y + 1}
	}

	trees := 0
	iterations := 0
	for ; err != types.ErrOffMap; err = b.Move(moveFunc) {
		iterations++
		if err == types.ErrTreeError {
			trees++
		}
	}

	fmt.Printf("trees in path: %d; iterations: %d; final x: %d\n", trees, iterations, b.Position.X)
}

func advance(p *types.Position) {
	p.X += 3
	p.Y++
}
