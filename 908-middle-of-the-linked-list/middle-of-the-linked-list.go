/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
    approach: slow and fast ptrs
    - if there are n nodes in the LL
    - and we have 2 ptrs; slow moving by 1 speed and fast moving by 2x speed
    - fast ptr will reach end in n/2 time
    - if fast ptr is moving at 2x speed, and fast has reached end
    - then slow which was moving at half of fast's speed
    - then slow will be at the middle of the LL

    time = o(n)
    space = o(1)
*/
func middleNode(head *ListNode) *ListNode {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}