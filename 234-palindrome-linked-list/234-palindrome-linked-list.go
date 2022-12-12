/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


/*
    approach : 1
    put the entire LL in a array
    and then check if the array is a palindrome or not ( two pointers from - 1 from start and 1 from end and move them against each other )
    time: o(n)
    space: o(n)
    
    approach : 2
    Reverse 2nd half of the LL
    and compare 2 LL starting from the heads
    time: o(n)
    space: o(1)
    
*/
func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true
    }
    
    var slowPrev *ListNode
    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        slowPrev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    
    head2 := slowPrev.Next
    slowPrev.Next = nil
    
    
    current := head2
    var prev *ListNode
    for current != nil {
        next := current.Next
        current.Next = prev
        prev = current
        current = next
    }
    head2 = prev
    
    
    for head != nil && head2 != nil {
        if head.Val != head2.Val {
            return false
        }
        head = head.Next
        head2 = head2.Next
    }
    return true
}