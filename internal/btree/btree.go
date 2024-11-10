package btree

type Node struct {
	keys		[]int // array of keys
	children	[]*Node // array of child pointers
	isLeaf		bool // check if leaf
}

type BTree struct {
	root		*Node
	minDegree	int 
}

// Insert a node
func (tree *BTree) Insert(key int) {
	if tree.root == nil {
		tree.root = &Node{isLeaf: true}
		tree.root.keys = append(tree.root.keys, key)
	} else {

	}
}