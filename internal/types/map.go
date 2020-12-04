package types

import (
	"errors"
)

var (
	// ErrTreeError indicates a tree found at position
	ErrTreeError = errors.New("tree encountered at position")
	// ErrOffMap indicates when a position is off the map
	ErrOffMap = errors.New("position off the map")
)

// NewMap creates a map from inputs
func NewMap(trees []Position, yLimit, copyFactor int) *Map {
	return &Map{trees: trees, yLimit: yLimit, copyFactor: copyFactor}
}

// Map is a position structure for locating trees
type Map struct {
	trees      []Position
	yLimit     int
	copyFactor int
}

// ReadPosition is a utility that returns errors for cases of off map or trees
func (m *Map) ReadPosition(p Position) error {
	// return error if Y value is off the map limit
	if p.Y > m.yLimit {
		return ErrOffMap
	}

	// adjust for copy factor
	for p.X > m.copyFactor {
		p.X = p.X - m.copyFactor
	}

	// loop through all trees on the map and return an error if one is
	// encountered
	for _, t := range m.trees {
		if p.X == t.X && p.Y == t.Y {
			return ErrTreeError
		}
	}

	return nil
}

// Toboggan holds location information and movement details for a mode
// of transportation.
type Toboggan struct {
	Position Position
	Map      *Map
}

// Move provides a utility to move the toboggan from one position to another.
// An error will be returned if the new position cannot be satisfied by the map.
func (t *Toboggan) Move(f PathFunc) error {
	t.Position = f(t.Position)
	return t.Map.ReadPosition(t.Position)
}

// Position is an abstract position indicator for a map
type Position struct {
	X, Y int
}

// PathFunc is an abstract function type to pass to the Toboggan
type PathFunc func(Position) Position
