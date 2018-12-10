package data

import (
	"errors"
	"fmt"
)

var (
	errNotExist       = errors.New("index is not existed")
	errTreeNil        = errors.New("tree is null")
	errTreeIndexExist = errors.New("tree index is existed")
)

type Node struct {
	lchild *Node
	rchild *Node
	height int //树的深度
	index  int
	data   CacheData
}

type CacheData struct {
	key   string
	value string
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
	return node.height
}

// 左旋转
//
//    node  BF = 2
//       \
//         prchild     ----->       prchild    BF = 1
//           \                        /   \
//           pprchild               node  pprchild
func llRotation(node *Node) *Node {
	prchild := node.rchild
	node.rchild = prchild.lchild
	prchild.lchild = node
	//更新节点 node 的高度
	node.height = max(getHeight(node.lchild), getHeight(node.rchild)) + 1
	//更新新父节点高度
	prchild.height = max(getHeight(prchild.lchild), getHeight(prchild.rchild)) + 1
	return prchild
}

// 右旋转
//             node  BF = -2
//              /
//         plchild     ----->       plchild    BF = 1
//            /                        /   \
//        pplchild                lchild   node

func rrRotation(node *Node) *Node {
	plchild := node.lchild
	node.lchild = plchild.rchild
	plchild.rchild = node
	node.height = max(getHeight(node.lchild), getHeight(node.rchild)) + 1
	plchild.height = max(getHeight(node), getHeight(plchild.lchild)) + 1
	return plchild
}

// 先左转再右转
//          node                  node
//         /            左          /     右
//      node1         ---->    node2     --->         node2
//          \                   /                     /   \
//          node2s           node1                 node1  node
func lrRotation(node *Node) *Node {
	plchild := llRotation(node.lchild) //左旋转
	node.lchild = plchild
	return rrRotation(node)

}

// 先右转再左转
//       node                  node
//          \          右         \         左
//          node1    ---->       node2     --->      node2
//          /                       \                /   \
//        node2                    node1           node  node1
func rlRotation(node *Node) *Node {
	prchild := rrRotation(node.rchild)
	node.rchild = prchild
	node.rchild = prchild
	return llRotation(node)
}

//处理节点高度问题
func handleBF(node *Node) *Node {
	if getHeight(node.lchild)-getHeight(node.rchild) == 2 {
		if getHeight(node.lchild.lchild)-getHeight(node.lchild.rchild) > 0 { //RR
			node = rrRotation(node)
		} else {
			node = lrRotation(node)
		}
	} else if getHeight(node.lchild)-getHeight(node.rchild) == -2 {
		if getHeight(node.rchild.lchild)-getHeight(node.rchild.rchild) < 0 { //LL
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
		if err := Midtraverse(node.lchild, handle); err != nil { //处理左子树
			return err
		}
		if err := Midtraverse(node.rchild, handle); err != nil { //处理右子树
			return err
		}
	}
	return nil
}

//插入节点 ---> 依次向上递归，调整树平衡
func Add(node *Node, index int, data CacheData) (*Node, error) {
	if node == nil {
		return &Node{lchild: nil, rchild: nil, index: index, data: data, height: 1}, nil
	}
	if node.index > index {
		node.lchild, _ = Add(node.lchild, index, data)
		node = handleBF(node)
	} else if node.index < index {
		node.rchild, _ = Add(node.rchild, index, data)
		node = handleBF(node)
	} else {
		return nil, errTreeIndexExist
	}
	node.height = max(getHeight(node.lchild), getHeight(node.rchild)) + 1
	return node, nil
}

//删除指定index节点
//查找节点 ---> 删除节点 ----> 调整树结构
//删除节点时既要遵循二叉搜索树的定义又要符合二叉平衡树的要求   ---> 重点处理删除节点的拥有左右子树的情况
func Delete(node *Node, index int) (*Node, error) {
	if node == nil {
		return nil, errNotExist
	}
	if node.index == index { //找到对应节点
		//如果没有左子树或者右子树 --->直接返回nil
		if node.lchild == nil && node.rchild == nil {
			return nil, nil
		} else if node.lchild == nil || node.rchild == nil { //若只存在左子树或者右子树
			if node.lchild != nil {
				return node.lchild, nil
			} else {
				return node.rchild, nil
			}
		} else { //左右子树都存在
			//查找前驱，替换当前节点,然后再进行依次删除  ---> 节点删除后，前驱替换当前节点 ---> 需遍历到最后，调整平衡度
			var n *Node
			//前驱
			n = node.lchild
			for {
				if n.rchild == nil {
					break
				}
				n = n.rchild
			}
			//
			n.data, node.data = node.data, n.data
			n.index, node.index = node.index, n.index
			node.lchild, _ = Delete(node.lchild, n.index)
		}
	} else if node.index > index {
		node.lchild, _ = Delete(node.lchild, index)
	} else { //node.index < index
		node.rchild, _ = Delete(node.rchild, index)
	}
	//删除节点后节点高度
	node.height = max(getHeight(node.lchild), getHeight(node.rchild)) + 1
	//调整树的平衡度
	node = handleBF(node)
	return node, nil
}

//查找并返回节点
func Modify(node *Node, index int, data CacheData) (*Node, error) {
	for {
		if node == nil {
			return nil, errNotExist
		}
		if index == node.index { //查找到index节点
			node.data = data
			return node, nil
		} else if index > node.index {
			node = node.rchild
		} else {
			node = node.lchild
		}
	}
}

//查找并返回节点
func Get(node *Node, index int) (*Node, error) {
	for {
		if node == nil {
			return nil, errNotExist
		}
		if index == node.index { //查找到index节点
			return node, nil
		} else if index > node.index {
			node = node.rchild
		} else {
			node = node.lchild
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
//	// 搜索index=4的节点
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
//	fmt.Println("before delete node index:4\n")
//	Midtraverse(tree, f)
//	tree, _ = Delete(tree, 4)
//	fmt.Println("after delete node index:4\n")
//	Midtraverse(tree, f)
//	fmt.Println("Modify index 5")
//	Modify(tree, 5, CacheData{"oooo", "iiiii"})
//	Midtraverse(tree, f)
//
//}
