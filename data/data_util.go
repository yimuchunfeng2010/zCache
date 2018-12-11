package data

import (
	"errors"
	)

var (
	errNotExist       = errors.New("Index is not existed")
	errTreeNil        = errors.New("tree is null")
	errTreeIndexExist = errors.New("tree Index is existed")
)

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

func max(data1 int, data2 int) int {
	if data1 > data2 {
		return data1
	}
	return data2
}

func getHeight(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

// 左旋转
//
//    node  BF = 2
//       \
//         pRchild     ----->       pRchild    BF = 1
//           \                        /   \
//           ppRchild               node  ppRchild
func llRotation(node *Node) *Node {
	pRchild := node.Rchild
	node.Rchild = pRchild.Lchild
	pRchild.Lchild = node
	//更新节点 node 的高度
	node.Height = max(getHeight(node.Lchild), getHeight(node.Rchild)) + 1
	//更新新父节点高度
	pRchild.Height = max(getHeight(pRchild.Lchild), getHeight(pRchild.Rchild)) + 1
	return pRchild
}

// 右旋转
//             node  BF = -2
//              /
//         pLchild     ----->       pLchild    BF = 1
//            /                        /   \
//        ppLchild                Lchild   node

func rrRotation(node *Node) *Node {
	pLchild := node.Lchild
	node.Lchild = pLchild.Rchild
	pLchild.Rchild = node
	node.Height = max(getHeight(node.Lchild), getHeight(node.Rchild)) + 1
	pLchild.Height = max(getHeight(node), getHeight(pLchild.Lchild)) + 1
	return pLchild
}

// 先左转再右转
//          node                  node
//         /            左          /     右
//      node1         ---->    node2     --->         node2
//          \                   /                     /   \
//          node2s           node1                 node1  node
func lrRotation(node *Node) *Node {
	pLchild := llRotation(node.Lchild) //左旋转
	node.Lchild = pLchild
	return rrRotation(node)

}

// 先右转再左转
//       node                  node
//          \          右         \         左
//          node1    ---->       node2     --->      node2
//          /                       \                /   \
//        node2                    node1           node  node1
func rlRotation(node *Node) *Node {
	pRchild := rrRotation(node.Rchild)
	node.Rchild = pRchild
	node.Rchild = pRchild
	return llRotation(node)
}

//处理节点高度问题
func handleBF(node *Node) *Node {
	if getHeight(node.Lchild)-getHeight(node.Rchild) == 2 {
		if getHeight(node.Lchild.Lchild)-getHeight(node.Lchild.Rchild) > 0 { //RR
			node = rrRotation(node)
		} else {
			node = lrRotation(node)
		}
	} else if getHeight(node.Lchild)-getHeight(node.Rchild) == -2 {
		if getHeight(node.Rchild.Lchild)-getHeight(node.Rchild.Rchild) < 0 { //LL
			node = llRotation(node)
		} else {
			node = rlRotation(node)
		}
	}
	return node
}

//中序遍历树，并根据钩子函数处理数据
func Midtraverse(node *Node, handle func(interface{}) error) error {
	if node == nil {
		return nil
	} else {
		if err := handle(node); err != nil {
			return err
		}
		if err := Midtraverse(node.Lchild, handle); err != nil { //处理左子树
			return err
		}
		if err := Midtraverse(node.Rchild, handle); err != nil { //处理右子树
			return err
		}
	}
	return nil
}

//插入节点 ---> 依次向上递归，调整树平衡
func Add(node *Node, Index string, data CacheData) (*Node, error) {
	if node == nil {
		return &Node{Lchild: nil, Rchild: nil, Index: Index, Data: data, Height: 1}, nil
	}
	if node.Index > Index {
		node.Lchild, _ = Add(node.Lchild, Index, data)
		node = handleBF(node)
	} else if node.Index < Index {
		node.Rchild, _ = Add(node.Rchild, Index, data)
		node = handleBF(node)
	} else {
		return nil, errTreeIndexExist
	}
	node.Height = max(getHeight(node.Lchild), getHeight(node.Rchild)) + 1
	return node, nil
}

//删除指定Index节点
//查找节点 ---> 删除节点 ----> 调整树结构
//删除节点时既要遵循二叉搜索树的定义又要符合二叉平衡树的要求   ---> 重点处理删除节点的拥有左右子树的情况
func Delete(node *Node, Index string) (*Node, error) {
	if node == nil {
		return nil, errNotExist
	}
	if node.Index == Index { //找到对应节点
		//如果没有左子树或者右子树 --->直接返回nil
		if node.Lchild == nil && node.Rchild == nil {
			return nil, nil
		} else if node.Lchild == nil || node.Rchild == nil { //若只存在左子树或者右子树
			if node.Lchild != nil {
				return node.Lchild, nil
			} else {
				return node.Rchild, nil
			}
		} else { //左右子树都存在
			//查找前驱，替换当前节点,然后再进行依次删除  ---> 节点删除后，前驱替换当前节点 ---> 需遍历到最后，调整平衡度
			var n *Node
			//前驱
			n = node.Lchild
			for {
				if n.Rchild == nil {
					break
				}
				n = n.Rchild
			}
			//
			n.Data, node.Data = node.Data, n.Data
			n.Index, node.Index = node.Index, n.Index
			node.Lchild, _ = Delete(node.Lchild, n.Index)
		}
	} else if node.Index > Index {
		node.Lchild, _ = Delete(node.Lchild, Index)
	} else { //node.Index < Index
		node.Rchild, _ = Delete(node.Rchild, Index)
	}
	//删除节点后节点高度
	node.Height = max(getHeight(node.Lchild), getHeight(node.Rchild)) + 1
	//调整树的平衡度
	node = handleBF(node)
	return node, nil
}

//查找并返回节点
func Modify(node *Node, Index string, data CacheData) (*Node, error) {
	for {
		if node == nil {
			return nil, errNotExist
		}
		if Index == node.Index { //查找到Index节点
			node.Data = data
			return node, nil
		} else if Index > node.Index {
			node = node.Rchild
		} else {
			node = node.Lchild
		}
	}
}

//查找并返回节点
func Get(node *Node, Index string) (*Node, error) {
	for {
		if node == nil {
			return nil, errNotExist
		}
		if Index == node.Index { //查找到Index节点
			return node, nil
		} else if Index > node.Index {
			node = node.Rchild
		} else {
			node = node.Lchild
		}
	}
}
//
////test
//func main() {
//	//打印匿名函数
//	f := func(node interface{}) error {
//		fmt.Println(node)
//		return nil
//	}
//	// 插入测试数据
//	tree, _ := Add(nil, 3, CacheData{"aaaa", "bbbb"})
//	tree, _ = Add(tree, 4, CacheData{"aaaa", "cccc"})
//	tree, _ = Add(tree, 5, CacheData{"aaaa", "bbbb"})
//	tree, _ = Add(tree, 7, CacheData{"aaaa", "bbbb"})
//	tree, _ = Add(tree, 6, CacheData{"aaaa", "bbbb"})
//	tree, _ = Add(tree, 15, CacheData{"aaaa", "bbbb"})
//	fmt.Println("Midtravese\n")
//	if err := Midtraverse(tree, f); err != nil {
//		fmt.Printf("Midtraverse failed err:%v\n", err)
//	}
//	// 搜索Index=4的节点
//	fmt.Println("\ntest Get in the tree")
//	node, err := Get(tree, 4)
//	if err != nil {
//		fmt.Printf("err = %s\n", err)
//	} else {
//		fmt.Printf("in the tree ,found the node:%v\n", node)
//	}
//
//	//测试删除节点
//	fmt.Println("\ntest Delete the node in the tree")
//	fmt.Println("before delete node Index:4\n")
//	Midtraverse(tree, f)
//	tree, _ = Delete(tree, 4)
//	fmt.Println("after delete node Index:4\n")
//	Midtraverse(tree, f)
//	fmt.Println("Modify Index 5")
//	Modify(tree, 5, CacheData{"oooo", "iiiii"})
//	Midtraverse(tree, f)
//
//}

