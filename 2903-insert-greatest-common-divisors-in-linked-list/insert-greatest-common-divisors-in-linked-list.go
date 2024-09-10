/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {return head}
    var prev *ListNode
    curr := head
    for curr != nil {
        next := curr.Next
        if prev != nil {
            // we have a,b
            newNode := &ListNode{Val: gcd(prev.Val, curr.Val)}
            prev.Next = newNode
            newNode.Next = curr
        }
        prev = curr
        curr = next        
    }    
    return head
}

func gcd(a, b int) int {
    // base
    if a == b {return a}

    // logic
    if a > b {
        return gcd(a-b, b)
    } else {
        return gcd(a, b-a)
    }
}