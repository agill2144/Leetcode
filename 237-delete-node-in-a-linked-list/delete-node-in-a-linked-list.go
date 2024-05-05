/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
    if node== nil {
		return
	}
	var prev *ListNode
	for node.Next != nil {
		nextVal := node.Next.Val
		node.Val = nextVal
		prev = node
		node =node.Next
	}
	prev.Next = nil
}