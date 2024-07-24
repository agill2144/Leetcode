/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil || k == 0 {return head}
    // ogHead := head
    size := 1
    curr := head
    for curr != nil && curr.Next != nil {
        size++
        curr = curr.Next
    }
    k %= size
    if k == size || k == 0 {return head}
    curr.Next = head
    var prev *ListNode
    count := 0
    curr = head
    for count != size-k {
        prev = curr
        curr = curr.Next
        count++
    }
    prev.Next = nil
    return curr

}

// rotate array intuition from: https://leetcode.com/problems/rotate-array/
// func rotateRight(head *ListNode, k int) *ListNode {
//     if head == nil || head.Next == nil || k == 0  {return head}
//     size := 0
//     curr := head
//     for curr != nil {
//         size++
//         curr = curr.Next
//     }
//     k %= size
//     if k == 0 {return head}
        
//     // 1. reverse the whole LL
//     head = reverse(head)
    
//     // 2. split the LL into 2 chunks ( k nodes, and remaining nodes )
//     i := 1
//     curr = head
//     for i != k {
//         i++
//         curr = curr.Next
//     }
//     firstChunkHead := head
//     firstChunkTail := head
//     secondChunkHead := curr.Next
//     curr.Next = nil
    
//     // 3. reverse the first chunk
//     firstChunkHead = reverse(firstChunkHead)
    
//     // 4. reverse the second chunk
//     secondChunkHead = reverse(secondChunkHead)
//     firstChunkTail.Next = secondChunkHead
//     return firstChunkHead
// }

// func reverse(head *ListNode) *ListNode {
//     var prev *ListNode
//     curr := head
//     for curr != nil {
//         next := curr.Next
//         curr.Next = prev
//         prev = curr
//         curr = next
//     }
//     return prev
// }