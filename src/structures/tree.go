package structures

/*BinaryTreeNode
@Description: Binary tree
*/
type BinaryTreeNode struct {
	Val        int
	LeftChild  *BinaryTreeNode
	RightChild *BinaryTreeNode
}

/*MultiTreeNode
@Description: Multi tree
*/
type MultiTreeNode struct {
	Val      int
	Children *[]MultiTreeNode
}
