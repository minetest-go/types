package types

import (
	"fmt"
	"math"
)

type Pos [3]int

// zero position
var PosZero = Pos{0, 0, 0}

// size of a mapblock
var MapBlockSize = Pos{16, 16, 16}

func NewPos(x, y, z int) *Pos {
	return &Pos{x, y, z}
}

func NewPosFromIndex(i int) *Pos {
	x := i % 16
	i /= 16
	y := i % 16
	i /= 16
	z := i % 16
	return NewPos(x, y, z)
}

func SortPos(p1, p2 *Pos) (*Pos, *Pos) {
	return &Pos{
			min(p1[0], p2[0]),
			min(p1[1], p2[1]),
			min(p1[2], p2[2]),
		}, &Pos{
			max(p1[0], p2[0]),
			max(p1[1], p2[1]),
			max(p1[2], p2[2]),
		}
}

func (p *Pos) X() int { return p[0] }
func (p *Pos) Y() int { return p[1] }
func (p *Pos) Z() int { return p[2] }

func (p *Pos) String() string {
	return fmt.Sprintf("Pos{%d,%d,%d}", p.X(), p.Y(), p.Z())
}

func (p1 *Pos) Add(p2 *Pos) *Pos {
	return &Pos{
		p1[0] + p2[0],
		p1[1] + p2[1],
		p1[2] + p2[2],
	}
}

func (p1 *Pos) Subtract(p2 *Pos) *Pos {
	return &Pos{
		p1[0] - p2[0],
		p1[1] - p2[1],
		p1[2] - p2[2],
	}
}

func (p *Pos) Divide(n float64) *Pos {
	return &Pos{
		int(math.Floor(float64(p[0]) / n)),
		int(math.Floor(float64(p[1]) / n)),
		int(math.Floor(float64(p[2]) / n)),
	}
}

func (p1 *Pos) Multiply(n int) *Pos {
	return &Pos{
		p1[0] * n,
		p1[1] * n,
		p1[2] * n,
	}
}

func (p *Pos) IsWithin(min, max *Pos) bool {
	return p[0] >= min[0] && p[0] <= max[0] &&
		p[1] >= min[1] && p[1] <= max[1] &&
		p[2] >= min[2] && p[2] <= max[2]
}

func (p *Pos) Index() int {
	return p[0] + (p[1] * 16) + (p[2] * 256)
}
