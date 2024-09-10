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

// greatest common divisior is highest such x that can divide both a and b evenly
// meaning a/x and b/x leaves remainder as 0
func gcd(a, b int) int {
    // base
    if a == b {return a}

    // logic
    if a > b {
        // early exit
        // eg; a = 10; b = 5
        if a % b == 0 {return b}
        return gcd(a-b, b)
    } else {
        // early exit
        // eg; a = 5; b = 10
        if b % a == 0 {return a}
        return gcd(a, b-a)
    }
}