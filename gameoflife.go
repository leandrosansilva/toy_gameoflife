package gameoflife

import (
	"errors"
	"log"
)

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

type Coord struct {
	X, Y int
}

type WorldMatrix [][]Cell

type World struct {
	// The current Matrix and the next generation one
	Matrices         [2]WorldMatrix
	UsingFirstMatrix bool
}

type Generator struct {
	World *World
	Rules []Rule
}

type RuleFilter func(coord Coord) bool

type RuleApplier func(neighbours NeighboursStates, live bool) bool

type Rule struct {
	Filter  RuleFilter
	Applier RuleApplier
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

func (this *World) GetLiveMatrix() *WorldMatrix {
	matrix, _ := this.GetMatrices()
	return matrix
}

func (this *World) GetDeadMatrix() *WorldMatrix {
	_, matrix := this.GetMatrices()
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
		return (*this.GetLiveMatrix())[x][y].IsLive(), nil
	}

	return false, errors.New("Invalid coord")
}

func (this *World) ActivateCell(coord Coord) error {
	if this.IsCoordValid(coord) {
		x, y := coord.Get()
		(*this.GetDeadMatrix())[x][y] = NewLiveCell()
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

		if (*this.GetLiveMatrix())[y][x].IsLive() {
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

func NewGenerator(world *World) Generator {
	rules := make([]Rule, 0)

	countNeighbours := func(neighbours NeighboursStates, expectedState CellState) int {
		// NOTE: this is similar to reduce(sum)
		count := 0

		for _, state := range neighbours {
			if state == expectedState {
				count++
			}
		}

		return count
	}

	rules = append(rules, NewRule(func(coord Coord) bool {
		// Applies to live cells
		live, _ := world.IsCellLive(coord)
		return live
	}, func(neighbours NeighboursStates, live bool) bool {
		nLiveNeighbours := countNeighbours(neighbours, ACTIVE_CELL)
		return nLiveNeighbours == 2 || nLiveNeighbours == 3
	}))

	rules = append(rules, NewRule(func(coord Coord) bool {
		// Applies to dead cells
		live, _ := world.IsCellLive(coord)
		return !live
	}, func(neighbours NeighboursStates, live bool) bool {
		return countNeighbours(neighbours, ACTIVE_CELL) == 3
	}))

	return Generator{world, rules}
}

func (this *Generator) Step() {
	_, inactiveMatrix := this.World.GetMatrices()

	this.World.ForEachCoordinate(func(coord Coord) {
		isLive, _ := this.World.IsCellLive(coord)

		for _, rule := range this.Rules {
			if rule.ApplyToCell(coord, isLive, this.World.GetCellNeighbours(coord)) {
				log.Printf("Spawning %s\n", coord)
				*(inactiveMatrix.RefToCell(coord)) = NewLiveCell()
				return
			}
		}

		log.Printf("Killing %s\n", coord)
		*(inactiveMatrix.RefToCell(coord)) = NewDeadCell()
	})

	this.World.SwitchMatrices()
}

func NewRule(filter RuleFilter, applier RuleApplier) Rule {
	return Rule{filter, applier}
}

func (this *Rule) ApplyToCell(coord Coord, live bool, neighbours NeighboursStates) bool {
	return this.Filter(coord) && this.Applier(neighbours, live)
}
