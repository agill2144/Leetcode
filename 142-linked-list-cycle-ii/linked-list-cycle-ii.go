/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {break}
    }
    if fast == nil || fast.Next == nil {return nil}
    p1 := head
    for p1 != slow {p1 = p1.Next; slow = slow.Next}
    return p1
}
// func detectCycle(head *ListNode) *ListNode {
//     set := map[*ListNode]bool{}
//     curr := head
//     for curr != nil {
//         if set[curr] {return curr}
//         set[curr] = true
//         curr = curr.Next
//     }
//     return nil
// }