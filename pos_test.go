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
