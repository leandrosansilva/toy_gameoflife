package gameoflife

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
