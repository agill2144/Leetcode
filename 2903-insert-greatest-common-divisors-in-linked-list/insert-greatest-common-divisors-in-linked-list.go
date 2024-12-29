/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {return head}
    curr := head
    for curr != nil {
        next := curr.Next
        if next != nil {
            curr.Next = &ListNode{Val:gcd(curr.Val, next.Val)}
            curr = curr.Next
            curr.Next = next
        }
        curr = next
    }
    return head
}

func gcd(a, b int) int {
    res := min(a, b)
    for res > 0 {
        if a % res == 0 && b % res == 0 {return res}
        res--
    }
    return 0
}