package types

type Area struct {
	Pos1 *Pos
	Pos2 *Pos
}

func NewArea(p1, p2 *Pos) *Area {
	p1, p2 = SortPos(p1, p2)
	return &Area{Pos1: p1, Pos2: p2}
}

// the inclusive size
func (a *Area) Size() *Pos {
	return a.Pos2.Subtract(a.Pos1).Add(NewPos(1, 1, 1))
}

func (a *Area) Volume() int {
	s := a.Size()
	return s.X() * s.Y() * s.Z()
}

func (a *Area) Corners() []*Pos {
	return []*Pos{
		{a.Pos1.X(), a.Pos1.Y(), a.Pos1.Z()},
		{a.Pos2.X(), a.Pos1.Y(), a.Pos1.Z()},
		{a.Pos1.X(), a.Pos2.Y(), a.Pos1.Z()},
		{a.Pos1.X(), a.Pos1.Y(), a.Pos2.Z()},

		{a.Pos2.X(), a.Pos2.Y(), a.Pos2.Z()},
		{a.Pos1.X(), a.Pos2.Y(), a.Pos2.Z()},
		{a.Pos2.X(), a.Pos1.Y(), a.Pos2.Z()},
		{a.Pos2.X(), a.Pos2.Y(), a.Pos1.Z()},
	}
}
