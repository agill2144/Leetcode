/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func hasCycle(head *ListNode) bool {
    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {return true}
    }
    return false
}

 // time = o(n)
 // space = o(n)
// func hasCycle(head *ListNode) bool {
//     m := map[*ListNode]struct{}{}
//     curr := head
//     for curr != nil {
//         if _, ok := m[curr]; ok {return true}
//         m[curr] = struct{}{}
//         curr = curr.Next
//     }
//     return false
// }