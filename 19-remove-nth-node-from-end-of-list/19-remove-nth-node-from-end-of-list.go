/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
    Approach 1:
    - Get the LL size
    - Then size-n
    - Then while currentPos is not at the diff of above, keep moving ( while saving a ref to prev node )
    - Once at the position, prev.Next = current.Next and drop current.Next
    
    time: o(n)
    space: o(1)
    
    
    Approach 2:
    - Use 2 pointers
    - Move 1st pointer n times
    - Then move both pointers until the first pointer is not nil
    - This will make sure that the distance between the 2 is n
    - and once the first pointer ( that was moved n times ), reaches nil, then that means the 2nd pointer is n distance behind
    - which is the node we want to remove
    - ( still have to maintain a ref prev node )
    
    time: o(n)
    space: o(1)

*/
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
    
//     if head == nil {
//         return head
//     }
    
//     size := 1
//     current := head
//     for current != nil {
//         size++
//         current = current.Next
//     }
    
    
//     desiredPos := size-n
//     currentPos := 1
//     c := head
//     var prev *ListNode
//     // 1-2-3-4-5
//     // 0 1 2 3
//     //       c
//     for currentPos != desiredPos {
//         prev = c
//         c = c.Next
//         currentPos++
        
//     }
//     if prev == nil {
//         // we want to remove the head
//         newHead := head.Next
//         head.Next = nil
//         head = newHead
//         return head
//     }
    
//     prev.Next = c.Next
//     c.Next = nil
    
//     return head
// }


func removeNthFromEnd(head *ListNode, n int) *ListNode {
    
    if head == nil {
        return head
    }
    p1 := head
    p2 := head
    for i := 0; i < n; i++ {
        p2 = p2.Next
    }
    var prev *ListNode
    for p2 != nil {
        prev = p1
        p1 = p1.Next
        p2 = p2.Next
    }
    if prev == nil {
        newHead := head.Next
        head.Next = nil
        head = newHead
        return head
    }
    prev.Next = p1.Next
    p1.Next = nil
    return head
}