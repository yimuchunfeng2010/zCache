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

type DataNode struct {
	Index int64
	Key  string
	Value   string
	Next *DataNode

}
type LogMsg struct {
	File string
	Time time.Time
	Msg string
}

type CoreInfo struct {
	KeyNum int
}


type LogInfoNode struct {
	Msg string
	Next *LogInfoNode
}