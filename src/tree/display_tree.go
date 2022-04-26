package tree

import (
	"fmt"
	"go_for_algorithm/src/structures"
)

func DisplayTree(root *structures.BinaryTreeNode) {
	front, rear := -1, -1
	queue := make([]*structures.BinaryTreeNode, 2000)
	rear++
	last := rear
	queue[rear] = root
	for front < rear {
		front++
		t := queue[front]
		fmt.Printf("%d ", t.Val)
		if t.LeftChild != nil {
			rear++
			queue[rear] = t.LeftChild
		}
		if t.RightChild != nil {
			rear++
			queue[rear] = t.RightChild
		}
		if front == last {
			last = rear
			fmt.Println()
		}
	}
}
