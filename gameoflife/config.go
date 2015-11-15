package gameoflife

import (
	"encoding/json"
)

type Config struct {
	Size struct {
		Height int
		Width  int
	}

	// a coordinate is an array with two elements
	Positions [][2]int
}

func ParseConfig(configContent string) (Config, error) {
	var config Config

	if err := json.Unmarshal([]byte(configContent), &config); err != nil {
		return Config{}, err
	}

	return config, nil
}
