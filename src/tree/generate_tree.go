package tree

import . "go_for_algorithm/src/structures"

func GenerateBinaryTree(values []int) (root *BinaryTreeNode) {
	nodes := make([]BinaryTreeNode, len(values))
	root = &nodes[0]
	for i := range nodes {
		nodes[i].Val = values[i]
		lChild, rChild := i*2+1, i*2+2
		if lChild < len(nodes) {
			nodes[i].LeftChild = &nodes[lChild]
		}
		if rChild < len(nodes) {
			nodes[i].RightChild = &nodes[rChild]
		}
	}
	return root
}
