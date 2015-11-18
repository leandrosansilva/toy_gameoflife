package gameoflife

import (
	"errors"
	"fmt"
	"strings"
)

type Importer struct {
}

func NewSpecieImporter() Importer {
	return Importer{}
}

func (this *Importer) ImportFromString(content string) (Specie, error) {
	charToLiveSpecieCell := func(c rune, line int) (int, error) {
		if c == '*' {
			return 1, nil
		}

		if c == '.' {
			return 0, nil
		}

		return 0, errors.New(fmt.Sprintf("Invalid char \"%c\" on line %d", c, line))
	}

	buildRow := func(lineNumber int, line string) ([]int, error) {
		var err error

		row := make([]int, len(line))

		for index, c := range line {
			row[index], err = charToLiveSpecieCell(c, lineNumber)

			if err != nil {
				return []int{}, err
			}
		}

		return row, nil
	}

	specieRows := make([][]int, 0)

	for lineNumber, line := range strings.Split(content, "\n") {
		trimmed := strings.Trim(line, " ")

		// Ignore comments and empty lines
		if len(trimmed) == 0 || trimmed[0] == '#' {
			continue
		}

		row, err := buildRow(lineNumber, trimmed)

		if err != nil {
			return Specie{}, err
		}

		specieRows = append(specieRows, row)
	}

	return NewSpecie(specieRows)
}
