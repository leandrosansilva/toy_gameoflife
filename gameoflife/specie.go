package gameoflife

import "errors"

type Specie [][]int

func NewSpecie(desc [][]int) (Specie, error) {
	if len(desc) == 0 || len(desc[0]) == 0 {
		return Specie{}, errors.New("Invalid specie")
	}

	return Specie(desc), nil
}

func (this *Specie) Size() (h, w int) {
	return len(*this), len((*this)[0])
}
