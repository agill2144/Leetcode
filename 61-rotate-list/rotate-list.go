/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// apply the same reversing strategy from: https://leetcode.com/problems/rotate-array/
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil || k == 0  {return head}
    size := 0
    curr := head
    for curr != nil {
        size++
        curr = curr.Next
    }
    k %= size
    if k == 0 {return head}
        
    // reverse the whole LL
    head = reverse(head)
    
    // split the LL into 2 chunks ( k nodes, and remaining nodes )
    i := 1
    curr = head
    for i < k {
        i++
        curr = curr.Next
    }
    firstChunkHead := head
    firstChunkTail := head
    secondChunkHead := curr.Next
    curr.Next = nil
    
    // reverse the first chunk
    firstChunkHead = reverse(firstChunkHead)
    
    // reverse the second chunk
    secondChunkHead = reverse(secondChunkHead)
    firstChunkTail.Next = secondChunkHead
    return firstChunkHead
}

func reverse(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}