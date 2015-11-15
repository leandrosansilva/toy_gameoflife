package gameoflife

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
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
	})

	Convey("Test Rules", t, func() {
		Convey("Always live regardless of neighbours", func() {
			rule := NewRule(func(coord Coord) bool {
				return true
			}, func(neighbours NeighboursStates, coord Coord, live bool) bool {
				return true
			})

			So(rule.ApplyToCell(NewCoord(0, 0), true, NeighboursStates{}), ShouldBeTrue)
		})

		Convey("Always die regardless of neighbours", func() {
			rule := NewRule(func(coord Coord) bool {
				return true
			}, func(neighbours NeighboursStates, coord Coord, live bool) bool {
				return false
			})

			So(rule.ApplyToCell(NewCoord(0, 0), true, NeighboursStates{}), ShouldBeFalse)
		})

		Convey("Die because rule does not apply", func() {
			rule := NewRule(func(coord Coord) bool {
				x, y := coord.Get()
				return x == 0 && y == 0
			}, func(neighbours NeighboursStates, coord Coord, live bool) bool {
				return true
			})

			So(rule.ApplyToCell(NewCoord(1, 1), true, NeighboursStates{}), ShouldBeFalse)
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

			world.ForEachCoordinate(func(coord Coord) {
				live, err := world.IsCellLive(coord)
				So(err, ShouldEqual, nil)
				So(live, ShouldBeFalse)
			})
		})

		Convey("Cell with a single neighbour dies", func() {
			world, _ := NewWorld(3, 3)
			world.ActivateCell(NewCoord(1, 0))
			world.ActivateCell(NewCoord(1, 1))

			generator := NewGenerator(&world)
			generator.Step()

			world.ForEachCoordinate(func(coord Coord) {
				live, err := world.IsCellLive(coord)
				So(err, ShouldEqual, nil)
				So(live, ShouldBeFalse)
			})
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
}
