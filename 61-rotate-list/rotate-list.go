/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil || k == 0 {return head}
    size := 1
    curr := head
    for curr.Next != nil {
        curr = curr.Next
        size++
    }
    k %= size
    if k == 0 {return head}
    curr.Next = head
    dist := size-k
    count := 1
    curr = head
    for count < dist {
        count++
        curr = curr.Next
    }
    head = curr.Next
    curr.Next = nil
    return head
}