package gameoflife

import (
	"errors"
)

type Cell struct {
	Active bool
}

type Coordinate struct {
	X, Y int
}

type World struct {
	Matrix [][]Cell
}

func NewActiveCell() Cell {
	return Cell{true}
}

func NewInactiveCell() Cell {
	return Cell{false}
}

func (this *Cell) IsActive() bool {
	return this.Active
}

func NewCoord(x, y int) Coordinate {
	return Coordinate{x, y}
}

func (this *Coordinate) Get() (x, y int) {
	return this.X, this.Y
}

func NewWorld(h, w int) (World, error) {
	if h > 0 && w > 0 {
		matrix := make([][]Cell, w)

		for index, _ := range matrix {
			matrix[index] = make([]Cell, h)
		}

		return World{matrix}, nil
	}

	return World{}, errors.New("Impossible world")
}

func (this *World) IsCoordValid(coord Coordinate) bool {
	x, y := coord.Get()

	return x >= 0 && x < len(this.Matrix) && y >= 0 && y < len(this.Matrix[0])

}

func (this *World) IsCellActive(coord Coordinate) (bool, error) {
	if this.IsCoordValid(coord) {
		x, y := coord.Get()
		return this.Matrix[x][y].IsActive(), nil
	}

	return false, errors.New("Invalid coord")
}

func (this *World) ActivateCell(coord Coordinate) error {
	if this.IsCoordValid(coord) {
		x, y := coord.Get()
		this.Matrix[x][y] = NewActiveCell()
		return nil
	}

	return errors.New("Invalid coord")
}
