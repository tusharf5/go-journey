package bst

type BTree struct {
	root *Node
}

type Node struct {
	data       int
	leftChild  *Node
	rightChild *Node
}

// A binary search tree is a tree where left child node is smaller than the node and the right child node is greater than it
func New() *BTree {

	root := &BTree{}

	return root
}

func (d *BTree) Add(data int) {
	node := &Node{
		data: data,
	}
	if d.root == nil {
		d.root = node
	} else {
		addNode(node, d.root)
	}

}

func (d *Node) Exist() bool {
	return d != nil
}

func addNode(newNode *Node, currentNode *Node) {
	if newNode.data < currentNode.data {
		if currentNode.leftChild.Exist() {
			addNode(newNode, currentNode.leftChild)
		} else {
			currentNode.leftChild = newNode
		}
	} else if newNode.data > currentNode.data {
		if currentNode.rightChild.Exist() {
			addNode(newNode, currentNode.rightChild)
		} else {
			currentNode.rightChild = newNode
		}
	}
}

func (d *BTree) Traverse() {
	arr := make([]int, 0)

	arr = append(arr, d.root.data)

	// curr

}
