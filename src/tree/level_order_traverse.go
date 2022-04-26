package tree

import . "go_for_algorithm/src/structures"

/*
	@Title LevelOrderTraverse
 	@Description: level-order traverse of binary tree
 	@Author WuChengPei 2022-04-26 15:44:07
 	@Param root: root node of tree
 	@Return res: traverse output sequence
*/
func LevelOrderTraverse(root *BinaryTreeNode) (res []int) {
	front, rear := -1, -1
	queue := make([]*BinaryTreeNode, 2000)
	rear++
	queue[rear] = root
	for front < rear {
		front++
		t := queue[front]
		res = append(res, t.Val)
		if t.LeftChild != nil {
			rear++
			queue[rear] = t.LeftChild
		}
		if t.RightChild != nil {
			rear++
			queue[rear] = t.RightChild
		}
	}
	return
}
