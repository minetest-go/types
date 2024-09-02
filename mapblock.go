package types

import "fmt"

type MapBlock struct {
	Size         int            `json:"size"`
	Version      byte           `json:"version"`
	Underground  bool           `json:"underground"`
	AirOnly      bool           `json:"air_only"`
	Timestamp    uint32         `json:"timestamp"`
	Mapdata      *MapData       `json:"mapdata"`
	Metadata     *Metadata      `json:"metadata"`
	BlockMapping map[int]string `json:"blockmapping"`
}

func NewMapblock() *MapBlock {
	mb := MapBlock{}
	mb.Metadata = NewMetadata()
	mb.BlockMapping = make(map[int]string)
	return &mb
}

// returns true if the mapblock is empty (air-only)
func (mb *MapBlock) IsEmpty() bool {
	return len(mb.BlockMapping) == 0
}

func (mb *MapBlock) GetNodeId(p *Pos) int {
	return mb.Mapdata.ContentId[p.Index()]
}

func (mb *MapBlock) GetParam2(p *Pos) int {
	return mb.Mapdata.Param2[p.Index()]
}

func (mb *MapBlock) GetNodeName(p *Pos) string {
	id := mb.GetNodeId(p)
	return mb.BlockMapping[id]
}

func (mb *MapBlock) GetNode(p *Pos) (*Node, error) {
	i := p.Index()

	if i > len(mb.Mapdata.ContentId) {
		return nil, fmt.Errorf("unexpected index, got %d, len: %d, pos: %s", i, len(mb.Mapdata.ContentId), p)
	}

	return &Node{
		Pos:    p,
		Name:   mb.BlockMapping[mb.Mapdata.ContentId[i]],
		Param1: mb.Mapdata.Param1[i],
		Param2: mb.Mapdata.Param2[i],
	}, nil
}
