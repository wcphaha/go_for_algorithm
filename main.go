package main

import (
	"go_for_algorithm/src/tree"
)

func main() {
	nums := []int{1, 2, 3, 4}
	t := tree.GenerateBinaryTree(nums)
	//fmt.Println(t.Val, t.LeftChild, t.RightChild)
	tree.DisplayTree(t)
}
