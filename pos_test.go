package types_test

import (
	"testing"

	"github.com/minetest-go/types"
	"github.com/stretchr/testify/assert"
)

func TestPos(t *testing.T) {
	p := types.NewPos(0, 0, 0)
	assert.NotNil(t, p)

	p2 := p.Add(types.NewPos(10, 20, 30))
	assert.NotNil(t, p2)
	assert.Equal(t, 10, p2.X())
	assert.Equal(t, 20, p2.Y())
	assert.Equal(t, 30, p2.Z())
}

func TestPosIndex(t *testing.T) {
	pos := types.NewPos(5, 8, 12)

	i := pos.Index()
	pos = types.NewPosFromIndex(i)

	assert.Equal(t, pos.X(), 5)
	assert.Equal(t, pos.Y(), 8)
	assert.Equal(t, pos.Z(), 12)
}

func TestSortPos(t *testing.T) {
	p1 := &types.Pos{1, 2, 3}
	p2 := &types.Pos{3, 2, 1}

	p1, p2 = types.SortPos(p1, p2)
	assert.Equal(t, 1, p1[0])
	assert.Equal(t, 2, p1[1])
	assert.Equal(t, 1, p1[2])
	assert.Equal(t, 3, p2[0])
	assert.Equal(t, 2, p2[1])
	assert.Equal(t, 3, p2[2])
}

func TestPosIsWithin(t *testing.T) {
	p1 := &types.Pos{1, 2, 3}
	p2 := &types.Pos{10, 20, 30}

	assert.True(t, types.NewPos(1, 2, 3).IsWithin(p1, p2))
	assert.True(t, types.NewPos(10, 2, 3).IsWithin(p1, p2))
	assert.False(t, types.NewPos(0, 2, 3).IsWithin(p1, p2))
}
