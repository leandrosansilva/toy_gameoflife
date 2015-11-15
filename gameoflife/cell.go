package gameoflife

type CellState int

const (
	// Not used by the cell, but informs there
	// is no neighbour
	INVALID_NEIGHBOUR CellState = iota

	ACTIVE_CELL   CellState = iota
	INACTIVE_CELL CellState = iota
)

type NeighboursStates [8]CellState

type Cell struct {
	State CellState
}

func NewLiveCell() Cell {
	return Cell{ACTIVE_CELL}
}

func NewDeadCell() Cell {
	return Cell{INACTIVE_CELL}
}

func (this *Cell) IsLive() bool {
	return this.State == ACTIVE_CELL
}
