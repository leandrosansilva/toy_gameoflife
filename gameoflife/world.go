package gameoflife

import "errors"

type WorldMatrix [][]Cell

type World struct {
	// The current Matrix and the next generation one
	Matrices         [2]WorldMatrix
	UsingFirstMatrix bool
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
	return &(*this)[y][x]
}

func NewWorld(h, w int) (World, error) {
	if h > 0 && w > 0 {
		return World{[2]WorldMatrix{CreateMatrix(h, w), CreateMatrix(h, w)}, true}, nil
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
	x, y := coord.Get()
	return x >= 0 && x < len(this.Matrices[0]) && y >= 0 && y < len(this.Matrices[0][0])
}

func (this *World) IsCellLive(coord Coord) (bool, error) {
	if this.IsCoordValid(coord) {
		x, y := coord.Get()
		return (*this.GetActiveMatrix())[x][y].IsLive(), nil
	}

	return false, errors.New("Invalid coord")
}

func (this *World) ActivateCell(coord Coord) error {
	if this.IsCoordValid(coord) {
		x, y := coord.Get()
		(*this.GetActiveMatrix())[x][y] = NewLiveCell()

		return nil
	}

	return errors.New("Invalid coord")
}

func (this *World) ForEachCoordinate(f func(coord Coord)) {
	for iY, _ := range this.Matrices[0] {
		for iX, _ := range this.Matrices[0][iY] {
			f(NewCoord(iX, iY))
		}
	}
}

func (this *World) GetCellState(coord Coord) CellState {
	if this.IsCoordValid(coord) {
		x, y := coord.Get()

		if (*this.GetActiveMatrix())[y][x].IsLive() {
			return ACTIVE_CELL
		}

		return INACTIVE_CELL
	}

	return INVALID_NEIGHBOUR
}

func (this *World) GetCellNeighbours(coord Coord) NeighboursStates {
	return NeighboursStates{
		this.GetCellState(coord.NorthWest()),
		this.GetCellState(coord.North()),
		this.GetCellState(coord.NorthEast()),
		this.GetCellState(coord.East()),
		this.GetCellState(coord.SouthEast()),
		this.GetCellState(coord.South()),
		this.GetCellState(coord.SouthWest()),
		this.GetCellState(coord.West()),
	}
}

func (this *World) Size() (h, w int) {
	return len(this.Matrices[0][0]), len(this.Matrices[0])
}
