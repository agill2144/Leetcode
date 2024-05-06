/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// monostack ngr, process top pattern
// time = o(n)
// space = o(n)
func removeNodes(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {return head}
    st := []*ListNode{}
    curr := head
    for curr != nil {
        next := curr.Next
        for len(st) != 0 && curr.Val > st[len(st)-1].Val {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            top.Next = nil
            if len(st) != 0 {
                st[len(st)-1].Next = curr
            }
        }
        st = append(st, curr)
        curr = next
    }
    return st[0]
}