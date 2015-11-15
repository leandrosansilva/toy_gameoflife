package gameoflife

type Generator struct {
	World *World
	Rules []Rule
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
		// Applies to dead cells
		live := world.GetActiveMatrix().RefToCell(coord).IsLive()
		return !live
	}, func(neighbours NeighboursStates, coord Coord, live bool) bool {
		return countNeighbours(neighbours, ACTIVE_CELL) == 3
	}))

	rules = append(rules, NewRule(func(coord Coord) bool {
		// Applies to live cells
		live := world.GetActiveMatrix().RefToCell(coord).IsLive()
		return live
	}, func(neighbours NeighboursStates, coord Coord, live bool) bool {
		nLiveNeighbours := countNeighbours(neighbours, ACTIVE_CELL)
		return nLiveNeighbours == 2 || nLiveNeighbours == 3
	}))

	return Generator{world, rules}
}

func (this *Generator) Step() {
	activeMatrix, inactiveMatrix := this.World.GetMatrices()

	this.World.ForEachCoordinate(func(coord Coord) {
		isLive := activeMatrix.RefToCell(coord).IsLive()

		for _, rule := range this.Rules {
			neighbours := this.World.GetCellNeighbours(coord)

			if rule.ApplyToCell(coord, isLive, neighbours) {
				*(inactiveMatrix.RefToCell(coord)) = NewLiveCell()
				return
			}
		}

		*(inactiveMatrix.RefToCell(coord)) = NewDeadCell()
	})

	this.World.SwitchMatrices()
}
