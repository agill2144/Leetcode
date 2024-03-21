/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
    approach: classic
    - point to prev
    - prev at the end becomes the new 
    time = o(n)
    space = o(1)
*/
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {return head}
    curr := head
    var prev *ListNode
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}