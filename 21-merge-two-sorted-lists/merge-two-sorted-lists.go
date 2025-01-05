/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{Val: 0}
    tail := dummy
    l1, l2 := list1, list2
    for l1 != nil || l2 != nil {
        l1Val := math.MaxInt64
        l2Val := math.MaxInt64
        if l1 != nil {l1Val = l1.Val}
        if l2 != nil {l2Val = l2.Val}
        if l1Val < l2Val {
            tail.Next = l1
            l1 = l1.Next
        } else {
            tail.Next = l2
            l2 = l2.Next
        }
        tail = tail.Next
    }
    return dummy.Next
}