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

func merge2Lists(list1, list2 *ListNode) *ListNode {
    if list1 == nil {return list2}
    if list2 == nil {return list1}
    dummy := &ListNode{}
    tail := dummy
    for list1 != nil || list2 != nil {
        l1Val := math.MaxInt64
        l2Val := math.MaxInt64
        if list1 != nil {l1Val = list1.Val}
        if list2 != nil {l2Val = list2.Val}
        if l1Val < l2Val {
            tail.Next = list1
            tail = tail.Next
            list1 = list1.Next
        } else {
            tail.Next = list2
            tail = tail.Next
            list2 = list2.Next
        }
    }
    return dummy.Next
}