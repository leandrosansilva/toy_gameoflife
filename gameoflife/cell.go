package gameoflife

type CellState uint8

const (
	// Not used by the cell, but informs there
	// is no neighbour
	INVALID_NEIGHBOUR CellState = iota

	ACTIVE_CELL   CellState = iota
	INACTIVE_CELL CellState = iota
)

type NeighboursStates [8]CellState

type Cell bool

func NewLiveCell() Cell {
	return Cell(true)
}

func NewDeadCell() Cell {
	return Cell(false)
}

func (this *Cell) IsLive() bool {
	return bool(*this)
}
