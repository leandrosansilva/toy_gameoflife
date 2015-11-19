package gameoflife

import "errors"

type WorldMatrix struct {
	/*
	   [1,2]
	   [3,4] is internally [1,2,3,4]
	*/
	Array []Cell
	Width int
}

type World struct {
	// The current Matrix and the next generation one
	Matrices                     [2]WorldMatrix
	UsingFirstMatrix             bool
	NeighbourCoordTransformation CoordTransformation
}

func CreateMatrix(h, w int) WorldMatrix {
	return WorldMatrix{make([]Cell, w*h), w}
}

func (this *WorldMatrix) RefToCell(coord Coord) *Cell {
	x, y := coord.Get()
	return &(this.Array[y*this.Width+x])
}

func NewGenericWorld(h, w int, transformation CoordTransformation) (World, error) {
	if h > 0 && w > 0 {
		return World{[2]WorldMatrix{CreateMatrix(h, w), CreateMatrix(h, w)}, true, transformation}, nil
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

func (this *World) GetMatrices() (live, inactive *WorldMatrix) {
	if this.UsingFirstMatrix {
		return &this.Matrices[0], &this.Matrices[1]
	}

	return &this.Matrices[1], &this.Matrices[0]
}

func (this *World) GetActiveMatrix() *WorldMatrix {
	matrix, _ := this.GetMatrices()
	return matrix
}

func (this *World) SwapMatrices() {
	this.UsingFirstMatrix = !this.UsingFirstMatrix
}

func (this *World) IsCoordValid(coord Coord) bool {
	h, w := this.Size()
	x, y := coord.Get()

	return x >= 0 && x < w && y >= 0 && y < h
}

func (this *World) IsCellLive(coord Coord) (bool, error) {
	if this.IsCoordValid(coord) {
		return this.GetActiveMatrix().RefToCell(coord).IsLive(), nil
	}

	return false, errors.New("Invalid coord")
}

func (this *World) ActivateCell(coord Coord) error {
	if !this.IsCoordValid(coord) {
		return errors.New("Invalid coord")
	}

	(*this.GetActiveMatrix().RefToCell(coord)) = NewLiveCell()
	return nil
}

func (this *World) ForEachCoordinate(f func(coord Coord)) {
	l, w := len(this.Matrices[0].Array), this.Matrices[0].Width

	for i := 0; i < l; i++ {
		x, y := i%w, i/w
		f(NewCoord(x, y))
	}
}

func (this *World) GetCellState(coord Coord) CellState {
	if !this.IsCoordValid(coord) {
		return INVALID_NEIGHBOUR
	}

	if this.GetActiveMatrix().RefToCell(coord).IsLive() {
		return ACTIVE_CELL
	}

	return INACTIVE_CELL
}

func (this *World) GetCellNeighbours(coord Coord) NeighboursStates {
	return NeighboursStates{
		this.GetCellState(this.NeighbourCoordTransformation(coord.NorthWest())),
		this.GetCellState(this.NeighbourCoordTransformation(coord.North())),
		this.GetCellState(this.NeighbourCoordTransformation(coord.NorthEast())),
		this.GetCellState(this.NeighbourCoordTransformation(coord.East())),
		this.GetCellState(this.NeighbourCoordTransformation(coord.SouthEast())),
		this.GetCellState(this.NeighbourCoordTransformation(coord.South())),
		this.GetCellState(this.NeighbourCoordTransformation(coord.SouthWest())),
		this.GetCellState(this.NeighbourCoordTransformation(coord.West())),
	}
}

func (this *World) Size() (h, w int) {
	return len(this.Matrices[0].Array) / this.Matrices[0].Width, this.Matrices[0].Width
}
