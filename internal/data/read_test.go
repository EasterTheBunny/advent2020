package data

import (
	"bytes"
	"strings"
	"testing"

	"github.com/easterthebunny/advent2020/internal/types"
)

func TestReadPasswordData(t *testing.T) {
	data := []byte(`1-3 a: abcde`)
	r := bytes.NewReader(data)

	entries, _ := ReadPasswordData(r)
	if len(*entries) != 1 {
		t.Errorf("unexpected number of entries %d; expected %d", len(*entries), 1)
	}
}

func TestReadMapData(t *testing.T) {
	data := `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

	t.Run("ParseEncodeMapLine", func(t *testing.T) {
		line := `#...#...#..`
		trees, _ := parseEncodedMapLine(4, line)
		expected := []types.Position{
			{X: 0, Y: 4},
			{X: 4, Y: 4},
			{X: 8, Y: 4}}

		for i, r := range trees {
			if r.X != expected[i].X || r.Y != expected[i].Y {
				t.Errorf("unexpected tree position '%d, %d'; expected '%d, %d'", r.X, r.Y, expected[i].X, expected[i].Y)
			}
		}
	})

	t.Run("PointsTest", func(t *testing.T) {
		type tst struct {
			p        types.Position
			expected error
		}

		tests := []tst{
			{p: types.Position{X: 5, Y: 4}, expected: types.ErrTreeError},
			{p: types.Position{X: 5, Y: 11}, expected: types.ErrOffMap},
			{p: types.Position{X: 7, Y: 3}, expected: nil},
			{p: types.Position{X: 32, Y: 6}, expected: types.ErrTreeError},
			{p: types.Position{X: 65, Y: 6}, expected: types.ErrTreeError},
			{p: types.Position{X: 15, Y: 5}, expected: types.ErrTreeError},
			{p: types.Position{X: 18, Y: 6}, expected: nil},
			{p: types.Position{X: 5, Y: 6}, expected: types.ErrTreeError}}

		r := strings.NewReader(data)
		m, err := ReadMapData(r)
		if err != nil {
			t.Error(err)
		}

		for i, test := range tests {
			err = m.ReadPosition(test.p)
			if err != test.expected {
				t.Errorf("test index %d: unexpected result %s; expected %s", i, err, test.expected)
			}
		}
	})

	t.Run("CountTest", func(t *testing.T) {
		r := strings.NewReader(data)
		m, err := ReadMapData(r)
		if err != nil {
			t.Error(err)
		}

		type tst struct {
			path     types.PathFunc
			expected int
		}

		tests := []tst{
			{
				path: func(p types.Position) types.Position {
					return types.Position{X: p.X + 2, Y: p.Y + 1}
				},
				expected: 1,
			},
			{
				path: func(p types.Position) types.Position {
					return types.Position{X: p.X + 3, Y: p.Y + 1}
				},
				expected: 7,
			},
			{
				path: func(p types.Position) types.Position {
					return types.Position{X: p.X + 4, Y: p.Y + 1}
				},
				expected: 2,
			},
			{
				path: func(p types.Position) types.Position {
					return types.Position{X: p.X + 5, Y: p.Y + 1}
				},
				expected: 3,
			},
		}

		for i, test := range tests {
			trees := 0
			err = nil

			b := &types.Toboggan{
				Position: types.Position{X: 0, Y: 0},
				Map:      m}

			for ; err != types.ErrOffMap; err = b.Move(test.path) {
				if err == types.ErrTreeError {
					trees++
				}
			}

			if trees != test.expected {
				t.Errorf("test index: %d; incorrect tree count %d; expected %d", i, trees, test.expected)
			}

		}
	})
}
