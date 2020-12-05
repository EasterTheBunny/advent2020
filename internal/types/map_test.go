package types

import (
	"testing"
)

func TestTobogganMove(t *testing.T) {
	data := []Position{
		{X: 2, Y: 0}, {X: 3, Y: 0},
		{X: 0, Y: 1}, {X: 4, Y: 1}, {X: 8, Y: 1},
		{X: 1, Y: 2}, {X: 6, Y: 2}, {X: 9, Y: 2},
		{X: 2, Y: 3}, {X: 4, Y: 3}, {X: 8, Y: 3}, {X: 10, Y: 3},
		{X: 1, Y: 4}, {X: 5, Y: 4}, {X: 6, Y: 4}, {X: 9, Y: 4},
		{X: 2, Y: 5}, {X: 4, Y: 5}, {X: 5, Y: 5},
		{X: 1, Y: 6}, {X: 3, Y: 6}, {X: 5, Y: 6}, {X: 10, Y: 6},
		{X: 1, Y: 7}, {X: 10, Y: 7},
		{X: 0, Y: 8}, {X: 2, Y: 8}, {X: 3, Y: 8}, {X: 7, Y: 8},
		{X: 0, Y: 9}, {X: 4, Y: 9}, {X: 5, Y: 9}, {X: 10, Y: 9},
		{X: 1, Y: 10}, {X: 4, Y: 10}, {X: 8, Y: 10}, {X: 10, Y: 10}}

	m := NewMap(data, 10, 10)
	b := &Toboggan{
		Position: Position{X: 0, Y: 0},
		Map:      m}

	moveFunc := func(p Position) Position {
		return Position{X: p.X + 3, Y: p.Y + 1}
	}

	trees := 0
	var err error
	for ; err != ErrOffMap; err = b.Move(moveFunc) {
		if err == ErrTreeError {
			trees++
		}
	}

	if trees != 7 {
		t.Errorf("unexpected tree count %d; expected %d", trees, 7)
	}
}
