package gameoflife

import (
	"strings"
)

type Printer struct {
	World *World
}

func NewPrinter(world *World) Printer {
	return Printer{world}
}

func (this *Printer) PrintHorizontalBorder() string {
	_, w := this.World.Size()
	return strings.Repeat("#", w+2) + "\n"
}

func (this *Printer) PrintLine(line, w int) string {
	var output string

	output += "#"

	for i := 0; i < w; i++ {
		if live, _ := this.World.IsCellLive(NewCoord(i, line)); live {
			output += "o"
			continue
		}

		output += " "
	}

	output += "#\n"

	return output
}

func (this *Printer) Print() string {
	var output string

	output += this.PrintHorizontalBorder()

	h, w := this.World.Size()

	for i := 0; i < h; i++ {
		output += this.PrintLine(i, w)
	}

	output += this.PrintHorizontalBorder()

	return output
}
