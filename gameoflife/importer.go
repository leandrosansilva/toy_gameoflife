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

// FIXME: this method is enoooormous and MUST be refactored
// It also should somehow support Life 1.06
func (this *Importer) ImportFromString(content string) (Specie, error) {
	charToLiveSpecieCell := func(c rune, line int) (int, error) {
		if c == '*' {
			return 1, nil
		}

		if c == 'O' {
			return 1, nil
		}

		if c == '.' {
			return 0, nil
		}

		return 0, errors.New(fmt.Sprintf("Invalid char \"%c\" on line %d", c, line))
	}

	buildRow := func(lineNumber, length int, line string) ([]int, error) {
		var err error

		row := make([]int, length)

		for index, c := range line {
			row[index], err = charToLiveSpecieCell(c, lineNumber)

			if err != nil {
				return []int{}, err
			}
		}

		return row, nil
	}

	filterValidContent := func(content string) []string {
		result := make([]string, 0)
		splitted := strings.Split(content, "\n")

		for _, line := range splitted {
			trimmed := strings.Trim(line, " ")

			// Ignore comments and empty lines
			if len(trimmed) == 0 || trimmed[0] == '#' {
				continue
			}

			result = append(result, trimmed)
		}

		return result
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	validContent := filterValidContent(content)

	width := func() int {
		m := 0

		for _, line := range validContent {
			m = max(m, len(line))
		}

		return m
	}()

	specieRows := make([][]int, 0)

	for lineNumber, line := range validContent {
		row, err := buildRow(lineNumber, width, line)

		if err != nil {
			return Specie{}, err
		}

		specieRows = append(specieRows, row)
	}

	return NewSpecie(specieRows)
}
