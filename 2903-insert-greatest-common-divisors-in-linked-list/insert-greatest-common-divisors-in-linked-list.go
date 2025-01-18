/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
    if head == nil {return head}
    curr := head
    for curr != nil && curr.Next != nil{
        next := curr.Next
        res := &ListNode{Val:gcd(curr.Val, next.Val)}
        curr.Next = res
        res.Next = next
        curr = next
    }
    return head
}

func gcd(n1, n2 int) int {
    x := min(n1,n2)
    for x > 0 {
        if n1%x == 0 && n2%x == 0 {return x}
        x--
    }
    return 0
}