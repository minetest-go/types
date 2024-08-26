package types_test

import (
	"testing"

	"github.com/minetest-go/types"
	"github.com/stretchr/testify/assert"
)

func TestArea(t *testing.T) {
	p1 := types.NewPos(1, 2, 3)
	p2 := types.NewPos(10, 20, 30)

	a := types.NewArea(p1, p2)
	s := a.Size()
	assert.Equal(t, types.NewPos(10, 19, 28), s)

	v := a.Volume()
	assert.Equal(t, 5320, v)

	visited := map[string]bool{}
	for _, p := range a.Corners() {
		assert.False(t, visited[p.String()])
		visited[p.String()] = true
	}
}
