package structures

/*LinkedList
@Description: linked list node
*/
type LinkedList struct {
	Val  int
	Pre  *LinkedList
	Next *LinkedList
}
