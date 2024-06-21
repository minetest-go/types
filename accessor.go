package types

import "image/color"

// returns the node at the given position, nil if no node found
type NodeAccessor func(pos *Pos) (*Node, error)

// resolves the node-name and param2 to a color, nil if no color-mapping found
type ColorResolver func(name string, param2 int) *color.RGBA
