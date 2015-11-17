package gameoflife

import (
	"bytes"
	"encoding/json"
	"time"
)

type Duration time.Duration

type Config struct {
	Size struct {
		Height int
		Width  int
	}

	GenerationDuration Duration

	RandomCells int

	Circular bool

	// a coordinate is an array with two elements
	Positions [][2]int

	Species map[string]Specie

	Population []struct {
		Specie   string
		Position Coord
	}
}

func (this *Duration) UnmarshalText(text []byte) error {
	var b bytes.Buffer

	b.Write(text)

	duration, err := time.ParseDuration(b.String())

	if err == nil {
		*this = Duration(duration)
	}

	return err
}

func ParseConfig(configContent string) (Config, error) {
	var config Config

	if err := json.Unmarshal([]byte(configContent), &config); err != nil {
		return Config{}, err
	}

	return config, nil
}
