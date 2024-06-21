package types

type Metadata struct {
	Inventories map[int]map[string]*Inventory `json:"inventories"`
	Pairs       map[int]map[string]string     `json:"pairs"`
}

func NewMetadata() *Metadata {
	md := Metadata{}
	md.Inventories = make(map[int]map[string]*Inventory)
	md.Pairs = make(map[int]map[string]string)
	return &md
}

func (md *Metadata) GetMetadata(p *Pos) map[string]string {
	return md.GetPairsMap(p.Index())
}

func (md *Metadata) GetPairsMap(pos int) map[string]string {
	pairsMap := md.Pairs[pos]
	if pairsMap == nil {
		pairsMap = make(map[string]string)
		md.Pairs[pos] = pairsMap
	}

	return pairsMap
}

func (md *Metadata) GetInventoryMap(index int) map[string]*Inventory {
	invMap := md.Inventories[index]
	if invMap == nil {
		invMap = make(map[string]*Inventory)
		md.Inventories[index] = invMap
	}

	return invMap
}

func (md *Metadata) GetInventoryMapAtPos(p *Pos) map[string]*Inventory {
	return md.GetInventoryMap(p.Index())
}

func (md *Metadata) GetInventory(index int, name string) *Inventory {
	m := md.GetInventoryMap(index)
	inv := m[name]
	if inv == nil {
		inv = &Inventory{}
		m[name] = inv
	}

	return inv
}
