package main

import (
	"errors"
	"flag"
	"fmt"
	. "github.com/leandrosansilva/toy_gameoflife/gameoflife"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type ImportedSpecies map[string]string

func (this *ImportedSpecies) Set(value string) error {
	if *this == nil {
		*this = make(ImportedSpecies)
	}

	s := strings.Split(value, "=")

	if len(s) != 2 {
		return errors.New(fmt.Sprintf("Could not parse option \"%s\"", value))
	}

	if _, found := (*this)[s[0]]; found {
		return errors.New(fmt.Sprintf("Cannot define imported life \"%s\" more than once", value))
	}

	(*this)[s[0]] = s[1]

	return nil
}

func (this *ImportedSpecies) String() string {
	return "[]"
}

func main() {
	var configFilename string
	var showHelp bool
	var importedSpecies ImportedSpecies

	flag.BoolVar(&showHelp, "help", false, "Show help message")
	flag.StringVar(&configFilename, "config", "config.json", "Configuration file path")
	flag.Var(&importedSpecies, "i", "List of lifename=filename for imported life")

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

	world, _ := func() (World, error) {
		if config.Circular {
			return NewCircularWorld(config.Size.Height, config.Size.Width)
		}

		return NewWorld(config.Size.Height, config.Size.Width)
	}()

	for _, position := range config.Positions {
		world.ActivateCell(position)
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < config.RandomCells; i++ {
		x := rand.Int() % config.Size.Width
		y := rand.Int() % config.Size.Height

		world.ActivateCell(NewCoord(x, y))
	}

	importer := NewSpecieImporter()

	if config.Species == nil {
		config.Species = make(map[string]Specie)
	}

	for lifeName, filename := range importedSpecies {
		fileContent, err := ioutil.ReadFile(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not open file %s: \"%s\"\n", filename, err)
			os.Exit(3)
		}

		config.Species[lifeName], err = importer.ImportFromString(string(fileContent))

		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not import from file %s: \"%s\"\n", filename, err)
			os.Exit(4)
		}
	}

	placer := NewLifePlacer(&world)

	for _, life := range config.Population {
		specie, found := config.Species[life.Specie]

		if !found {
			fmt.Fprintf(os.Stderr, "Invalid specie %s\n", life.Specie)
			os.Exit(1)
		}

		if err := placer.Place(specie, life.Position); err != nil {
			fmt.Fprintf(os.Stderr, "Could not insert %s in position %s: \"%s\"\n", life.Specie, life.Position, err)
			os.Exit(1)
		}
	}

	generator := NewGenerator(&world)

	printer := NewPrinter(&world)

	start := time.Now()

	// yes, config.Generations == 0 means infinite loop :-)
	for i := uint64(0); i < config.Generations || config.Generations == 0; i++ {
		fmt.Print("\033[2J")
		fmt.Print(printer.Print())
		time.Sleep(time.Duration(config.GenerationDuration))
		generator.Step()
	}

	elapsed := time.Since(start)

	fmt.Printf("Using %d steps has taken %s\n", config.Generations, elapsed)
}
