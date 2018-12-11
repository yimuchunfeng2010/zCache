package types
type Node struct {
	Lchild *Node
	Rchild *Node
	Height int //树的深度
	Key  string
	Value   string
}

