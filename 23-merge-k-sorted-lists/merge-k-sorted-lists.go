/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    merged := []int{}
    for i := 0; i < len(lists); i++ {
        curr := lists[i]
        for curr != nil {
            merged = append(merged, curr.Val)
            curr = curr.Next
        }
    }
    sort.Ints(merged)
    head := &ListNode{Val: 0}
    tail := head
    for i := 0; i < len(merged); i++ {
        tail.Next = &ListNode{Val: merged[i]}
        tail = tail.Next
    }
    return head.Next
}