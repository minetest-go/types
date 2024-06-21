package types

type Node struct {
	Pos    *Pos
	Name   string `json:"name"`
	Param1 int    `json:"param1"`
	Param2 int    `json:"param2"`
}
