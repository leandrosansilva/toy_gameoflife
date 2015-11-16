package main

import (
	"flag"
	"fmt"
	. "gameoflife/gameoflife"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	var configFilename string
	var showHelp bool

	flag.BoolVar(&showHelp, "help", false, "Show help message")
	flag.StringVar(&configFilename, "config", "config.json", "Configuration file path")

	flag.Parse()

	if showHelp {
		flag.Usage()
		os.Exit(2)
	}

	fileBytes, err := ioutil.ReadFile(configFilename)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open %s: %s\n", configFilename, err)
		os.Exit(1)
	}

	config, err := ParseConfig(string(fileBytes))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not parse file %s: %s\n", configFilename, err)
		os.Exit(1)
	}

	world, _ := NewWorld(config.Size.Height, config.Size.Width)

	for _, position := range config.Positions {
		world.ActivateCell(NewCoord(position[0], position[1]))
	}

	generator := NewGenerator(&world)

	printer := NewPrinter(&world)

	fmt.Print("\033[2J")

	for {
		fmt.Print(printer.Print())
		time.Sleep(time.Duration(config.GenerationDuration))
		fmt.Print("\033[2J")
		generator.Step()
	}
}
