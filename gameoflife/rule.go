package gameoflife

type RuleFilter func(coord Coord) bool

type RuleApplier func(neighbours NeighboursCoords, coord Coord) bool

type Rule struct {
	Filter  RuleFilter
	Applier RuleApplier
}

func NewRule(filter RuleFilter, applier RuleApplier) Rule {
	return Rule{filter, applier}
}

func (this *Rule) ApplyToCell(coord Coord, neighbours NeighboursCoords) bool {
	r := this.Applier(neighbours, coord)
	return r
}
