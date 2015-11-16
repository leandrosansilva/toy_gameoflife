package gameoflife

type Generator struct {
	World *World
	Rules []Rule
}

func CreateDefaultRules(world *World) []Rule {
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

	return []Rule{
		NewRule(func(coord Coord) bool {
			// Applies to dead cells
			return !world.GetActiveMatrix().RefToCell(coord).IsLive()
		}, func(neighbours NeighboursStates, coord Coord, live bool) bool {
			return countNeighbours(neighbours, ACTIVE_CELL) == 3
		}),

		NewRule(func(coord Coord) bool {
			// Applies to live cells
			return world.GetActiveMatrix().RefToCell(coord).IsLive()
		}, func(neighbours NeighboursStates, coord Coord, live bool) bool {
			nLiveNeighbours := countNeighbours(neighbours, ACTIVE_CELL)
			return nLiveNeighbours == 2 || nLiveNeighbours == 3
		}),
	}
}

func NewGenericGenerator(world *World, rules []Rule) Generator {
	return Generator{world, rules}
}

func NewGenerator(world *World) Generator {
	return NewGenericGenerator(world, CreateDefaultRules(world))
}

func (this *Generator) Step() {
	activeMatrix, inactiveMatrix := this.World.GetMatrices()

	this.World.ForEachCoordinate(func(coord Coord) {
		isLive := activeMatrix.RefToCell(coord).IsLive()

		for _, rule := range this.Rules {
			neighbours := this.World.GetCellNeighbours(coord)

			// NOTE: stops on the first rule which
			// applies to the cell
			if rule.ApplyToCell(coord, isLive, neighbours) {
				*(inactiveMatrix.RefToCell(coord)) = NewLiveCell()
				return
			}
		}

		*(inactiveMatrix.RefToCell(coord)) = NewDeadCell()
	})

	this.World.SwapMatrices()
}
