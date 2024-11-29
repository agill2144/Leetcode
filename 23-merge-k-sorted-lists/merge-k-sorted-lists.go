/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {return nil}
    pairIncr := 1
    iIncr := 2
    for pairIncr < len(lists) {
        for i := 0; i + pairIncr < len(lists); i += iIncr {
            lists[i] = merge2Lists(lists[i], lists[i+pairIncr])
        }
        iIncr *= 2
        pairIncr *= 2
    }
    return lists[0]
}

func merge2Lists(l1, l2 *ListNode) *ListNode {
    head := &ListNode{Val:0}
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