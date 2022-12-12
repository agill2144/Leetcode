/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// time: o(2n) or just o(n) 
// space: o(1)

func pairSum(head *ListNode) int {
    var prev *ListNode
    slow := head
    fast := head
    // reverse first half of the linked list
    for fast != nil && fast.Next != nil {
        slowNext := slow.Next
        fastNext := fast.Next.Next
        if fastNext == nil {
            slow.Next = nil
        }
        slow.Next = prev
        prev = slow
        slow = slowNext
        fast = fastNext
    }
    
    // break the ll into 2 halves
    max := 0
    
    // the first half is reversed
    for slow != nil {
        if slow.Val + prev.Val > max {
            max = slow.Val + prev.Val
        }
        slow = slow.Next
        prev = prev.Next
    }
    
    
    return max
}