/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    items := []int{}
    for i := 0; i < len(lists); i++ {
        curr := lists[i]
        for curr != nil {
            items = append(items, curr.Val)
            curr = curr.Next
        }
    }
    sort.Ints(items)
    head := &ListNode{Val: 0}
    tail := head
    for i := 0; i < len(items); i++ {
        tail.Next = &ListNode{Val: items[i]}
        tail = tail.Next
    }
    return head.Next
}