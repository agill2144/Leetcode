/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {return nil}
    head := lists[0]
    for i := 1; i < len(lists); i++ {
        head = merge2Lists(head, lists[i])
    }
    return head
}

func merge2Lists(l1, l2 *ListNode) *ListNode {
    if l1 == nil && l2 == nil {return nil}
    if l1 == nil && l2 != nil {return l2}
    if l1 != nil && l2 == nil {return l1}
    head := &ListNode{Val: 0}
    tail := head
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            tail.Next = l1
            tail = tail.Next
            l1 = l1.Next
        } else {
            tail.Next = l2
            tail = tail.Next
            l2 = l2.Next
        }
    }
    for l1 != nil {
        tail.Next = l1
        tail = tail.Next
        l1 = l1.Next
    }
    for l2 != nil {
        tail.Next = l2
        tail = tail.Next
        l2 = l2.Next
    }
    return head.Next
}