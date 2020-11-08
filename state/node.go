package state

import (
	"strings"
	"sync"
)

// Node is node level data of the tree ds
type Node struct {
	LevelType  string
	LevelValue string
	Webreq     *uint64
	Timespent  *uint64
	Children   map[string]*Node
	Mutex      *sync.Mutex
}

func GetKey(levelType, levelValue string) string {
	return strings.Trim(levelType, " ") + ":" + strings.Trim(levelValue, " ")
}

func GetInitialState() *Node {
	return &Node{LevelType: "global", LevelValue: "global", Webreq: new(uint64), Timespent: new(uint64), Children: make(map[string]*Node), Mutex: &sync.Mutex{}}
}
