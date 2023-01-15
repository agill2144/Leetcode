/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true
    }
    slow := head
    fast := head.Next
    
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    reversedHead := reverse(slow.Next)
    slow.Next = nil
    
    curr := head
    
    for reversedHead != nil {
        if curr.Val != reversedHead.Val {return false}
        curr = curr.Next
        reversedHead = reversedHead.Next
    }
    return true
}

func reverse(head *ListNode) *ListNode {
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