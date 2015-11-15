package gameoflife

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGameOfLife(t *testing.T) {

	Convey("Test Cells", t, func() {
		Convey("Test single active cell", func() {
			cell := NewActiveCell()
			So(cell.IsActive(), ShouldBeTrue)
		})

		Convey("Test single inactive cell", func() {
			cell := NewInactiveCell()
			So(cell.IsActive(), ShouldBeFalse)
		})
	})

	Convey("Test Coordinate", t, func() {
		Convey("Basic Coordinate", func() {
			coord := NewCoord(1, 3)
			x, y := coord.Get()
			So(x, ShouldEqual, 1)
			So(y, ShouldEqual, 3)
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
				_, err := world.IsCellActive(NewCoord(1, 0))
				So(err, ShouldResemble, errors.New("Invalid coord"))
			})

			Convey("Invalid Y", func() {
				_, err := world.IsCellActive(NewCoord(0, 1))
				So(err, ShouldResemble, errors.New("Invalid coord"))
			})

			Convey("Invalid X and Y", func() {
				_, err := world.IsCellActive(NewCoord(1, -3))
				So(err, ShouldResemble, errors.New("Invalid coord"))
			})
		})

		Convey("Activation of cell in trivial world", func() {
			world, _ := NewWorld(1, 1)

			Convey("Cell on 0,0 is inactive", func() {
				active, _ := world.IsCellActive(NewCoord(0, 0))
				So(active, ShouldBeFalse)
			})

			Convey("Fail to activate cell", func() {
				err := world.ActivateCell(NewCoord(1, 1))
				So(err, ShouldResemble, errors.New("Invalid coord"))
			})

			Convey("Active cell in 0,0", func() {
				err := world.ActivateCell(NewCoord(0, 0))
				So(err, ShouldEqual, nil)
				active, _ := world.IsCellActive(NewCoord(0, 0))
				So(active, ShouldBeTrue)
			})
		})
	})
}
