package gameoflife

type Generator struct {
	World *World
	Rules []Rule
}

func CreateDefaultRules(world *World) []Rule {
	matrix := world.GetActiveMatrix()

	countLiveNeighbours := func(neighbours NeighboursCoords) int {
		// NOTE: this is similar to reduce(sum)
		count := 0

		for _, coord := range neighbours {
			if matrix.IsLive(coord) {
				count++
			}
		}

		return count
	}

	return []Rule{
		NewRule(func(coord Coord) bool {
			// Applies to dead cells
			return !matrix.IsLive(coord)
		}, func(neighbours NeighboursCoords, coord Coord) bool {
			return countLiveNeighbours(neighbours) == 3
		}),

		NewRule(func(coord Coord) bool {
			// Applies to live cells
			return matrix.IsLive(coord)
		}, func(neighbours NeighboursCoords, coord Coord) bool {
			nLiveNeighbours := countLiveNeighbours(neighbours)
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
	activeMatrix := this.World.GetActiveMatrix()
	inactiveMatrix := this.World.GetInactiveMatrix()

	this.World.ForEachCoordinate(func(coord Coord) {
		neighbours := this.World.GetCellNeighboursCoords(coord)

		for _, n := range neighbours {
			if !activeMatrix.IsLive(n) && !inactiveMatrix.IsLive(n) {
				inactiveMatrix.SetCellState(n, false)
			}
		}

		inactiveMatrix.SetCellState(coord, func() bool {
			for _, rule := range this.Rules {
				if rule.Filter(coord) {
					return rule.ApplyToCell(coord, neighbours)
				}
			}

			return false
		}())
	})

	this.World.SwapMatrices()
}
