package main

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

type BST struct {
	Root *Node
}

func (bst *BST) Insert(key int) {
	bst.Root = insertNode(bst.Root, key)
}
func insertNode(root *Node, key int) *Node {
	if root == nil {
		return &Node{Key: key, Left: nil, Right: nil}
	}
	if key < root.Key {
		root.Left = insertNode(root.Left, key)
	} else if key > root.Key {
		root.Right = insertNode(root.Right, key)
	}
	return root
}

func main() {
	bst := &BST{}
	keys := []int{10, 7, 5, 8}
	for _, key := range keys {
		bst.Insert(key)
	}
}
