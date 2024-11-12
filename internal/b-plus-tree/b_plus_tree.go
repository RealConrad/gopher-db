package b_plus_tree

type Database struct {
	tables map[string]*Table // Key: Table names, Values: pointers to said table
}

type Table struct {
	name		string
	columns		[]string
	bPlusTree	*BPlusTree // Root of B+Tree for indexing
}

type BPlusTree struct {
	root	*BPlusTreeNode
	order	int // Defines the maximum number of children a node can have
}

type Record struct {
    values map[string]interface{}
}

type BPlusTreeNode struct {
	isLeaf		bool
	keys		[]interface{}
	children	[]*BPlusTreeNode // internal nodes
	records		[]*Record // Only for leaf nodes
	next		*BPlusTreeNode // linked list stucture in leaf nodes
}

func (tree *BPlusTree) Search(key interface{}) *Record {
	currentNode := tree.root;
	
	// Get to the correct leaf node
	for !currentNode.isLeaf {
		i := 0
		for i < len(currentNode.keys) && key >= currentNode.keys[i] {
			i++
		}
		currentNode = currentNode.children[i]
	}

	// At leaf node, find correct record
	for i, k := range(currentNode.keys) {
		if k == key {
			return currentNode.records[i]
		}
	}
	return nil // Node not found
}
