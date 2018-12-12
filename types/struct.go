package types

import "time"

type Node struct {
	Lchild *Node
	Rchild *Node
	Height int
	Key  string
	Value   string
}

type KeyValue struct{
	Key  string
	Value   string
}

type LogMsg struct {
	File string
	Time time.Time
	Msg string
}