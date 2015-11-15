package gameoflife

type RuleFilter func(coord Coord) bool

type RuleApplier func(neighbours NeighboursStates, coord Coord, live bool) bool

type Rule struct {
	Filter  RuleFilter
	Applier RuleApplier
}

func NewRule(filter RuleFilter, applier RuleApplier) Rule {
	return Rule{filter, applier}
}

func (this *Rule) ApplyToCell(coord Coord, live bool, neighbours NeighboursStates) bool {
	return this.Filter(coord) && this.Applier(neighbours, coord, live)
}
