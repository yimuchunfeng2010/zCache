package types
type Node struct {
	Lchild *Node
	Rchild *Node
	Height int //树的深度
	Index  string
	Data   CacheData
}

type CacheData struct {
	Key   string
	Value string
}
