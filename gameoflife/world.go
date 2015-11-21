package gameoflife

import (
	"errors"
)

type WorldMatrix map[Coord]bool

type World struct {
	// The current Matrix and the next generation one
	ActiveMatrix, InactiveMatrix WorldMatrix
	Height, Width                int
	NeighbourCoordTransformation CoordTransformation
}

func (this *WorldMatrix) IsLive(coord Coord) bool {
	state, found := (*this)[coord]
	return state && found
}

func (this *WorldMatrix) SetCellState(coord Coord, state bool) {
	(*this)[coord] = state
}

func CreateMatrix() WorldMatrix {
	return make(WorldMatrix)
}

func NewGenericWorld(h, w int, transformation CoordTransformation) (World, error) {
	if h > 0 && w > 0 {
		return World{CreateMatrix(), CreateMatrix(), h, w, transformation}, nil
	}

	return World{}, errors.New("Impossible world")
}

func NewWorld(h, w int) (World, error) {
	return NewGenericWorld(h, w, func(coord Coord) Coord {
		return coord
	})
}

func NewCircularWorld(h, w int) (World, error) {
	circulate := func(val, max int) int {
		if val < 0 {
			return max - 1
		}

		if val >= max {
			return 0
		}

		return val
	}

	return NewGenericWorld(h, w, func(coord Coord) Coord {
		x, y := coord.Get()
		return NewCoord(circulate(x, w), circulate(y, h))
	})
}

func (this *World) GetActiveMatrix() *WorldMatrix {
	return &this.ActiveMatrix
}

func (this *World) GetInactiveMatrix() *WorldMatrix {
	return &this.InactiveMatrix
}

func (this *World) SwapMatrices() {
	this.ActiveMatrix = this.InactiveMatrix
	this.InactiveMatrix = CreateMatrix()
}

func (this *World) IsCoordValid(coord Coord) bool {
	h, w := this.Size()
	x, y := coord.Get()

	return x >= 0 && x < w && y >= 0 && y < h
}

func (this *World) IsCellLive(coord Coord) (bool, error) {
	if this.IsCoordValid(coord) {
		return this.ActiveMatrix.IsLive(coord), nil
	}

	return false, errors.New("Invalid coord")
}

func (this *World) ActivateCell(coord Coord) error {
	if !this.IsCoordValid(coord) {
		return errors.New("Invalid coord")
	}

	for _, n := range this.GetCellNeighboursCoords(coord) {
		if !this.ActiveMatrix.IsLive(n) {
			this.ActiveMatrix.SetCellState(n, false)
		}
	}

	this.ActiveMatrix.SetCellState(coord, true)

	return nil
}

func (this *World) ForEachCoordinate(f func(Coord)) {
	for c, _ := range this.ActiveMatrix {
		f(c)
	}
}

func (this *World) GetCellNeighboursCoords(coord Coord) NeighboursCoords {
	validCoords := make(NeighboursCoords, 0, 8)

	for _, n := range []Coord{
		coord.NorthWest(),
		coord.North(),
		coord.NorthEast(),
		coord.East(),
		coord.SouthEast(),
		coord.South(),
		coord.SouthWest(),
		coord.West(),
	} {
		t := this.NeighbourCoordTransformation(n)
		if this.IsCoordValid(t) {
			validCoords = append(validCoords, t)
		}
	}

	return validCoords
}

func (this *World) GetCellLiveNeighboursCoords(coord Coord) NeighboursCoords {
	lives := make(NeighboursCoords, 0, 8)

	for _, n := range this.GetCellNeighboursCoords(coord) {
		if this.ActiveMatrix.IsLive(n) {
			lives = append(lives, n)
		}
	}

	return lives
}

func (this *World) Size() (h, w int) {
	return this.Height, this.Width
}
