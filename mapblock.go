package types

import "fmt"

type MapBlock struct {
	Size        int    `json:"size"`
	Version     byte   `json:"version"`
	Underground bool   `json:"underground"`
	AirOnly     bool   `json:"air_only"`
	Timestamp   uint32 `json:"timestamp"`
	ContentId   []int  `json:"contentid"`
	Param1      []int  `json:"param1"`
	Param2      []int  `json:"param2"`
	// index -> inventory-map
	Inventory map[int]map[string][]string `json:"inventory"`
	// index -> fields
	Fields map[int]map[string]string `json:"fields"`
	// nodeid -> nodename
	BlockMapping map[int]string `json:"blockmapping"`
}

func NewMapblock() *MapBlock {
	mb := MapBlock{}
	mb.Inventory = make(map[int]map[string][]string)
	mb.Fields = make(map[int]map[string]string)
	mb.BlockMapping = make(map[int]string)
	return &mb
}

// returns true if the mapblock is empty (air-only)
func (mb *MapBlock) IsEmpty() bool {
	return len(mb.BlockMapping) == 0
}

func (mb *MapBlock) GetNodeId(p *Pos) int {
	return mb.ContentId[p.Index()]
}

func (mb *MapBlock) GetParam2(p *Pos) int {
	return mb.Param2[p.Index()]
}

func (mb *MapBlock) GetNodeName(p *Pos) string {
	id := mb.GetNodeId(p)
	return mb.BlockMapping[id]
}

func (mb *MapBlock) GetNode(p *Pos) (*Node, error) {
	i := p.Index()

	if i > len(mb.ContentId) {
		return nil, fmt.Errorf("unexpected index, got %d, len: %d, pos: %s", i, len(mb.ContentId), p)
	}

	return &Node{
		Pos:    p,
		Name:   mb.BlockMapping[mb.ContentId[i]],
		Param1: mb.Param1[i],
		Param2: mb.Param2[i],
	}, nil
}
