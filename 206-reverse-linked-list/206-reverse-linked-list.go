
/*

Approach 1

space: o(n)
time: o(n) 

We can implement stack as a linked list.
So go thru each node and push them to a "stack using LL"
i.e push each node to the head/top
*/
// func reverseList(head *ListNode) *ListNode {
//     if head == nil || head.Next == nil {
//         return head
//     }
    
//     var stack *ListNode
//     current := head
//     for current != nil {
//         newHead := &ListNode{Val: current.Val}
//         if stack == nil {
//             stack = newHead
//         } else {
//             newHead.Next = stack
//             stack = newHead
//         }
//         current = current.Next
//     }
//     return stack
// }


/*

Approach 2

space: o(1)
time: o(n) 

reverse in place using 2 pointers ( current and prev)
(H)1->2->3->4
1<-2<-3<-4(H)
*/
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    var prev *ListNode
    current := head
    for current != nil {
        next := current.Next
        current.Next = prev
        prev = current
        current = next
    }
    return prev
}