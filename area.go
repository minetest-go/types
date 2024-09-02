package types

import "fmt"

type Area struct {
	Pos1 *Pos `json:"pos1"`
	Pos2 *Pos `json:"pos2"`
}

func NewArea(p1, p2 *Pos) *Area {
	p1, p2 = SortPos(p1, p2)
	return &Area{Pos1: p1, Pos2: p2}
}

func (a *Area) String() string {
	return fmt.Sprintf("Area{%v,%v}", a.Pos1, a.Pos2)
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

func (a *Area) Intersects(a2 *Area) bool {
	for _, p1 := range a.Corners() {
		if p1.IsWithin(a2.Pos1, a2.Pos2) {
			return true
		}
	}
	for _, p2 := range a2.Corners() {
		if p2.IsWithin(a.Pos1, a.Pos2) {
			return true
		}
	}
	return false
}

func (a *Area) Union(a2 *Area) *Area {
	if !a.Intersects(a2) {
		// no union possible
		return nil
	}

	return &Area{
		Pos1: NewPos(
			max(a.Pos1.X(), a2.Pos1.X()),
			max(a.Pos1.Y(), a2.Pos1.Y()),
			max(a.Pos1.Z(), a2.Pos1.Z()),
		),
		Pos2: NewPos(
			min(a.Pos2.X(), a2.Pos2.X()),
			min(a.Pos2.Y(), a2.Pos2.Y()),
			min(a.Pos2.Z(), a2.Pos2.Z()),
		),
	}
}
