package gameoflife

import (
	"encoding/json"
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestGameOfLife(t *testing.T) {

	Convey("Test Cells", t, func() {
		Convey("Test single live cell", func() {
			cell := NewLiveCell()
			So(cell.IsLive(), ShouldBeTrue)
		})

		Convey("Test single dead cell", func() {
			cell := NewDeadCell()
			So(cell.IsLive(), ShouldBeFalse)
		})
	})

	Convey("Test Coord", t, func() {
		Convey("Basic Coord", func() {
			coord := NewCoord(1, 3)
			x, y := coord.Get()
			So(x, ShouldEqual, 1)
			So(y, ShouldEqual, 3)
		})

		Convey("North", func() {
			coord := NewCoord(0, 0).North()
			x, y := coord.Get()
			So(x, ShouldEqual, 0)
			So(y, ShouldEqual, -1)
		})

		Convey("South", func() {
			coord := NewCoord(0, 0).South()
			x, y := coord.Get()
			So(x, ShouldEqual, 0)
			So(y, ShouldEqual, 1)
		})

		Convey("West", func() {
			coord := NewCoord(0, 0).West()
			x, y := coord.Get()
			So(x, ShouldEqual, -1)
			So(y, ShouldEqual, 0)
		})

		Convey("East", func() {
			coord := NewCoord(0, 0).East()
			x, y := coord.Get()
			So(x, ShouldEqual, 1)
			So(y, ShouldEqual, 0)
		})

		Convey("North-West", func() {
			coord := NewCoord(0, 0).NorthWest()
			x, y := coord.Get()
			So(x, ShouldEqual, -1)
			So(y, ShouldEqual, -1)
		})

		Convey("North-East", func() {
			coord := NewCoord(0, 0).NorthEast()
			x, y := coord.Get()
			So(x, ShouldEqual, 1)
			So(y, ShouldEqual, -1)
		})

		Convey("South-West", func() {
			coord := NewCoord(0, 0).SouthWest()
			x, y := coord.Get()
			So(x, ShouldEqual, -1)
			So(y, ShouldEqual, 1)
		})

		Convey("South-East", func() {
			coord := NewCoord(0, 0).SouthEast()
			x, y := coord.Get()
			So(x, ShouldEqual, 1)
			So(y, ShouldEqual, 1)
		})
	})

	Convey("Test World", t, func() {
		Convey("Impossible world", func() {
			_, err := NewWorld(0, 0)
			So(err, ShouldResemble, errors.New("Impossible world"))
		})

		Convey("Size", func() {
			world, _ := NewWorld(35, 42)
			h, w := world.Size()
			So(h, ShouldEqual, 35)
			So(w, ShouldEqual, 42)
		})

		Convey("10 columns and 2 rows world", func() {
			/*world, _ := NewWorld(10, 2)

			world.ForEachCoordinate(func(coord Coord) {
				So(world.IsCoordValid(coord), ShouldBeTrue)
				state := world.GetCellState(coord)
				So(state, ShouldEqual, INACTIVE_CELL)
			})*/
		})

		Convey("Invalid cell position", func() {
			world, err := NewWorld(1, 1)
			So(err, ShouldEqual, nil)

			Convey("Invalid X", func() {
				_, err := world.IsCellLive(NewCoord(1, 0))
				So(err, ShouldResemble, errors.New("Invalid coord"))
			})

			Convey("Invalid Y", func() {
				_, err := world.IsCellLive(NewCoord(0, 1))
				So(err, ShouldResemble, errors.New("Invalid coord"))
			})

			Convey("Invalid X and Y", func() {
				_, err := world.IsCellLive(NewCoord(1, -3))
				So(err, ShouldResemble, errors.New("Invalid coord"))
			})
		})

		Convey("Activation of cell in trivial world", func() {
			world, _ := NewWorld(1, 1)

			Convey("Cell on 0,0 is dead", func() {
				live, _ := world.IsCellLive(NewCoord(0, 0))
				So(live, ShouldBeFalse)
			})

			Convey("Fail to activate cell", func() {
				err := world.ActivateCell(NewCoord(1, 1))
				So(err, ShouldResemble, errors.New("Invalid coord"))
			})

			Convey("Live cell in 0,0", func() {
				err := world.ActivateCell(NewCoord(0, 0))
				So(err, ShouldEqual, nil)
				live, _ := world.IsCellLive(NewCoord(0, 0))
				So(live, ShouldBeTrue)
			})
		})

		Convey("Circular World", func() {
			world, _ := NewCircularWorld(3, 3)
			So(world.ActivateCell(NewCoord(0, 0)), ShouldEqual, nil)
			So(world.ActivateCell(NewCoord(0, 2)), ShouldEqual, nil)
			So(world.ActivateCell(NewCoord(2, 2)), ShouldEqual, nil)
			So(world.ActivateCell(NewCoord(2, 0)), ShouldEqual, nil)

			Convey("Cells on the edges have 3 neighbours", func() {
				Convey("0x0", func() {
					So(world.GetCellNeighbours(NewCoord(0, 0)), ShouldResemble, NeighboursStates{
						ACTIVE_CELL,
						ACTIVE_CELL,
						INACTIVE_CELL,
						INACTIVE_CELL,
						INACTIVE_CELL,
						INACTIVE_CELL,
						INACTIVE_CELL,
						ACTIVE_CELL,
					})
				})

				Convey("0x2", func() {
					So(world.GetCellNeighbours(NewCoord(0, 2)), ShouldResemble, NeighboursStates{
						INACTIVE_CELL,
						INACTIVE_CELL,
						INACTIVE_CELL,
						INACTIVE_CELL,
						INACTIVE_CELL,
						ACTIVE_CELL,
						ACTIVE_CELL,
						ACTIVE_CELL,
					})
				})

				Convey("2x2", func() {
					So(world.GetCellNeighbours(NewCoord(2, 2)), ShouldResemble, NeighboursStates{
						INACTIVE_CELL,
						INACTIVE_CELL,
						INACTIVE_CELL,
						ACTIVE_CELL,
						ACTIVE_CELL,
						ACTIVE_CELL,
						INACTIVE_CELL,
						INACTIVE_CELL,
					})
				})

				Convey("2x0", func() {
					So(world.GetCellNeighbours(NewCoord(2, 0)), ShouldResemble, NeighboursStates{
						INACTIVE_CELL,
						ACTIVE_CELL,
						ACTIVE_CELL,
						ACTIVE_CELL,
						INACTIVE_CELL,
						INACTIVE_CELL,
						INACTIVE_CELL,
						INACTIVE_CELL,
					})
				})
			})
		})
	})

	Convey("Test Rules", t, func() {
		Convey("Always live regardless of neighbours", func() {
			rule := NewRule(func(coord Coord) bool {
				return true
			}, func(neighbours NeighboursStates, coord Coord) bool {
				return true
			})

			So(rule.Filter(NewCoord(0, 0)) && rule.ApplyToCell(NewCoord(0, 0), NeighboursStates{}), ShouldBeTrue)
		})

		Convey("Always die regardless of neighbours", func() {
			rule := NewRule(func(coord Coord) bool {
				return true
			}, func(neighbours NeighboursStates, coord Coord) bool {
				return false
			})

			So(rule.Filter(NewCoord(0, 0)) && rule.ApplyToCell(NewCoord(0, 0), NeighboursStates{}), ShouldBeFalse)
		})

		Convey("Die because rule does not apply", func() {
			rule := NewRule(func(coord Coord) bool {
				x, y := coord.Get()
				return x == 0 && y == 0
			}, func(neighbours NeighboursStates, coord Coord) bool {
				return true
			})

			So(rule.Filter(NewCoord(1, 1)) && rule.ApplyToCell(NewCoord(1, 1), NeighboursStates{}), ShouldBeFalse)
		})
	})

	Convey("Test Neighbours", t, func() {
		Convey("All neighbours are invalid in the trivial world", func() {
			world, _ := NewWorld(1, 1)
			neighbours := world.GetCellNeighbours(NewCoord(0, 0))

			So(neighbours, ShouldEqual, NeighboursStates{
				INVALID_NEIGHBOUR,
				INVALID_NEIGHBOUR,
				INVALID_NEIGHBOUR,
				INVALID_NEIGHBOUR,
				INVALID_NEIGHBOUR,
				INVALID_NEIGHBOUR,
				INVALID_NEIGHBOUR,
				INVALID_NEIGHBOUR,
			})
		})

		Convey("Smallest square valid world", func() {
			// TODO: test more...
			world, _ := NewWorld(3, 3)

			So(world.GetCellNeighbours(NewCoord(1, 1)), ShouldEqual, NeighboursStates{
				INACTIVE_CELL,
				INACTIVE_CELL,
				INACTIVE_CELL,
				INACTIVE_CELL,
				INACTIVE_CELL,
				INACTIVE_CELL,
				INACTIVE_CELL,
				INACTIVE_CELL,
			})

			So(world.GetCellNeighbours(NewCoord(0, 0)), ShouldEqual, NeighboursStates{
				INVALID_NEIGHBOUR,
				INVALID_NEIGHBOUR,
				INVALID_NEIGHBOUR,
				INACTIVE_CELL,
				INACTIVE_CELL,
				INACTIVE_CELL,
				INVALID_NEIGHBOUR,
				INVALID_NEIGHBOUR,
			})
		})
	})

	Convey("Test Generator", t, func() {
		Convey("Cell alone dies", func() {
			world, _ := NewWorld(3, 3)
			world.ActivateCell(NewCoord(1, 1))

			generator := NewGenerator(&world)
			generator.Step()

			/*world.ForEachCoordinate(func(coord Coord) {
				live, err := world.IsCellLive(coord)
				So(err, ShouldEqual, nil)
				So(live, ShouldBeFalse)
			})*/
		})

		Convey("Cell with a single neighbour dies", func() {
			world, _ := NewWorld(3, 3)
			world.ActivateCell(NewCoord(1, 0))
			world.ActivateCell(NewCoord(1, 1))

			generator := NewGenerator(&world)
			generator.Step()

			/*world.ForEachCoordinate(func(coord Coord) {
				live, err := world.IsCellLive(coord)
				So(err, ShouldEqual, nil)
				So(live, ShouldBeFalse)
			})*/
		})

		Convey("Three cells inline rotate", func() {
			world, _ := NewWorld(3, 3)
			world.ActivateCell(NewCoord(1, 0))
			world.ActivateCell(NewCoord(1, 1))
			world.ActivateCell(NewCoord(1, 2))

			generator := NewGenerator(&world)
			generator.Step()

			Convey("0x0", func() {
				live, _ := world.IsCellLive(NewCoord(0, 0))
				So(live, ShouldBeFalse)
			})

			Convey("1x0", func() {
				live, _ := world.IsCellLive(NewCoord(1, 0))
				So(live, ShouldBeFalse)
			})

			Convey("2x0", func() {
				live, _ := world.IsCellLive(NewCoord(2, 0))
				So(live, ShouldBeFalse)
			})

			Convey("0x1", func() {
				live, _ := world.IsCellLive(NewCoord(0, 1))
				So(live, ShouldBeTrue)
			})

			Convey("1x1", func() {
				live, _ := world.IsCellLive(NewCoord(1, 1))
				So(live, ShouldBeTrue)
			})

			Convey("2x1", func() {
				live, _ := world.IsCellLive(NewCoord(2, 1))
				So(live, ShouldBeTrue)
			})

			Convey("0x2", func() {
				live, _ := world.IsCellLive(NewCoord(0, 2))
				So(live, ShouldBeFalse)
			})

			Convey("1x2", func() {
				live, _ := world.IsCellLive(NewCoord(1, 2))
				So(live, ShouldBeFalse)
			})

			Convey("2x2", func() {
				live, _ := world.IsCellLive(NewCoord(2, 2))
				So(live, ShouldBeFalse)
			})
		})
	})

	Convey("Test Output", t, func() {
		Convey("Trivial world", func() {
			world, _ := NewWorld(1, 1)
			printer := NewPrinter(&world)

			Convey("Empty", func() {
				So(printer.Print(), ShouldEqual, "###\n# #\n###\n")
			})

			Convey("One live cell", func() {
				world.ActivateCell(NewCoord(0, 0))
				So(printer.Print(), ShouldEqual, "###\n#o#\n###\n")
			})
		})
	})

	Convey("Test Config File", t, func() {
		Convey("Empty map", func() {
			jsonContent := `{"Size": {"Height": 10, "Width": 20}, "GenerationDuration": "500ms", "Positions": []}`
			config, err := ParseConfig(jsonContent)
			So(err, ShouldEqual, nil)
			So(config.Size.Height, ShouldEqual, 10)
			So(config.Size.Width, ShouldEqual, 20)
			So(config.GenerationDuration, ShouldEqual, time.Millisecond*500)
			So(config.RandomCells, ShouldEqual, 0)
			So(len(config.Positions), ShouldEqual, 0)
		})
	})

	Convey("Species", t, func() {
		Convey("Simplest", func() {
			simple, err := NewSpecie([][]int{{1}})
			So(err, ShouldEqual, nil)
			h, w := simple.Size()
			So(h, ShouldEqual, 1)
			So(w, ShouldEqual, 1)
			So(simple[0][0], ShouldEqual, 1)
		})

		Convey("Glider", func() {
			glider, err := NewSpecie([][]int{
				{0, 1, 0},
				{0, 0, 1},
				{1, 1, 1},
			})

			So(err, ShouldEqual, nil)
			h, w := glider.Size()
			So(h, ShouldEqual, 3)
			So(w, ShouldEqual, 3)
		})

		Convey("Letter L", func() {
			l, err := NewSpecie([][]int{
				{1, 0},
				{1, 0},
				{1, 1},
			})

			So(err, ShouldEqual, nil)
			h, w := l.Size()
			So(h, ShouldEqual, 3)
			So(w, ShouldEqual, 2)
		})
	})

	Convey("Create life", t, func() {
		Convey("Put Letter L on 4x4 world", func() {
			l, err := NewSpecie([][]int{
				{1, 0},
				{1, 0},
				{1, 1},
			})

			world, _ := NewWorld(4, 4)

			placer := NewLifePlacer(&world)

			Convey("Insert in invalid position", func() {
				err := placer.Place(l, NewCoord(3, 2))
				So(err, ShouldResemble, errors.New("Invalid position to form of life"))
			})

			err = placer.Place(l, NewCoord(2, 1))

			So(err, ShouldEqual, nil)

			Convey("Active cells", func() {
				for _, value := range []Coord{
					{2, 1}, {2, 2}, {2, 3}, {3, 3},
				} {
					active, err := world.IsCellLive(NewCoord(value[0], value[1]))
					So(err, ShouldEqual, nil)
					So(active, ShouldBeTrue)
				}
			})

			Convey("Inactive cells", func() {
				for _, value := range []Coord{
					{0, 0}, {1, 0}, {2, 0}, {3, 0},
					{0, 1}, {1, 1}, {3, 1},
					{0, 2}, {1, 2}, {3, 2},
					{0, 3}, {1, 3},
				} {
					active, err := world.IsCellLive(NewCoord(value[0], value[1]))
					So(err, ShouldEqual, nil)
					So(active, ShouldBeFalse)
				}
			})
		})
	})

	Convey("Parse config Specie", t, func() {
		Convey("Parse Glider", func() {
			configContent := "[[0,1,0],[0,0,1],[1,1,1]]"

			var specie Specie

			err := json.Unmarshal([]byte(configContent), &specie)

			So(err, ShouldEqual, nil)

			glider, _ := NewSpecie([][]int{
				{0, 1, 0},
				{0, 0, 1},
				{1, 1, 1},
			})

			So(glider, ShouldResemble, specie)
		})
	})

	Convey("Test Import .lif (1.05) files", t, func() {
		Convey("glider", func() {
			content := `#Life 1.05
#D This is a glider.
#N
#P -1 -1
.*
..*
***`
			importer := NewSpecieImporter()

			specie, err := importer.ImportFromString(content)

			So(err, ShouldEqual, nil)

			glider, _ := NewSpecie([][]int{
				{0, 1, 0},
				{0, 0, 1},
				{1, 1, 1},
			})

			So(specie, ShouldResemble, glider)

		})
	})

}
