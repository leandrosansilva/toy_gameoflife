package gameoflife

import "errors"

type LifePlacer struct {
	World *World
}

func NewLifePlacer(world *World) LifePlacer {
	return LifePlacer{world}
}

func (this *LifePlacer) Place(specie Specie, coord Coord) error {
	specieH, specieW := specie.Size()
	x, y := coord.Get()

	if !this.World.IsCoordValid(coord) || !this.World.IsCoordValid(NewCoord(x+specieW-1, y+specieH-1)) {
		return errors.New("Invalid position to form of life")
	}

	for itH := 0; itH < specieH; itH++ {
		for itW := 0; itW < specieW; itW++ {
			if specie[itH][itW] == 0 {
				continue
			}

			if err := this.World.ActivateCell(NewCoord(x+itW, y+itH)); err != nil {
				return err
			}
		}
	}

	return nil
}
