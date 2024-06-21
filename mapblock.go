package types

type MapBlock struct {
	Size         int            `json:"size"`
	Version      byte           `json:"version"`
	Underground  bool           `json:"underground"`
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
