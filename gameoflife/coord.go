package gameoflife

type Coord struct {
	X, Y int
}

type CoordTransformation func(Coord) Coord

func NewCoord(x, y int) Coord {
	return Coord{x, y}
}

func (this Coord) North() Coord {
	x, y := this.Get()
	return NewCoord(x, y-1)
}

func (this Coord) South() Coord {
	x, y := this.Get()
	return NewCoord(x, y+1)
}

func (this Coord) West() Coord {
	x, y := this.Get()
	return NewCoord(x-1, y)
}

func (this Coord) East() Coord {
	x, y := this.Get()
	return NewCoord(x+1, y)
}

func (this Coord) NorthWest() Coord {
	return this.North().West()
}

func (this Coord) NorthEast() Coord {
	return this.North().East()
}

func (this Coord) SouthWest() Coord {
	return this.South().West()
}

func (this Coord) SouthEast() Coord {
	return this.South().East()
}

func (this *Coord) Get() (x, y int) {
	return this.X, this.Y
}
