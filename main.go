package main

import (
	"fmt"
	. "gameoflife/gameoflife"
	"time"
)

func main() {
	world, _ := NewWorld(30, 30)
	world.ActivateCell(NewCoord(10, 8))
	world.ActivateCell(NewCoord(10, 9))
	world.ActivateCell(NewCoord(10, 10))

	generator := NewGenerator(&world)

	printer := NewPrinter(&world)

	fmt.Print("\033[2J")

	for {
		fmt.Print(printer.Print())
		time.Sleep(time.Second * 1)
		fmt.Print("\033[2J")
		generator.Step()
	}
}
