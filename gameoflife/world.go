package gameoflife

import "errors"

type WorldMatrix [][]Cell

type World struct {
	// The current Matrix and the next generation one
	Matrices                     [2]WorldMatrix
	UsingFirstMatrix             bool
	NeighbourCoordTransformation CoordTransformation
}

func CreateMatrix(h, w int) WorldMatrix {
	matrix := make(WorldMatrix, w)

	for iw, _ := range matrix {
		matrix[iw] = make([]Cell, h)

		for ih, _ := range matrix[iw] {
			matrix[iw][ih] = NewDeadCell()
		}
	}

	return matrix
}

func (this *WorldMatrix) RefToCell(coord Coord) *Cell {
	x, y := coord.Get()
	return &(*this)[x][y]
}

func NewWorld(h, w int) (World, error) {
	neighbourCoordTransformation := func(coord Coord) Coord {
		return coord
	}

	if h > 0 && w > 0 {
		return World{[2]WorldMatrix{CreateMatrix(h, w), CreateMatrix(h, w)}, true, neighbourCoordTransformation}, nil
	}

	return World{}, errors.New("Impossible world")
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

func (this *World) SwitchMatrices() {
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
	if this.IsCoordValid(coord) {
		(*this.GetActiveMatrix().RefToCell(coord)) = NewLiveCell()
		return nil
	}

	return errors.New("Invalid coord")
}

func (this *World) ForEachCoordinate(f func(coord Coord)) {
	h, w := this.Size()

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			f(NewCoord(x, y))
		}
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
	return len(this.Matrices[0][0]), len(this.Matrices[0])
}
