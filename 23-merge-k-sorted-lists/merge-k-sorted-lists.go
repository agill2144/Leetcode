/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {return nil}
    var dfs func(left, right int) *ListNode
    dfs = func(left, right int) *ListNode {
        // base
        if left >= right {
            if left == right {return lists[left]}
            return nil
        }

        // logic
        mid := left + (right-left)/2
        l1 := dfs(left, mid)
        l2 := dfs(mid+1, right)
        return merge2Lists(l1, l2)
    }
    return dfs(0, len(lists)-1)
}

func merge2Lists(l1, l2 *ListNode) *ListNode {
    if l1 == nil {return l2}
    if l2 == nil {return l1}
    head := &ListNode{Val: 0}
    tail := head
    for l1 != nil || l2 != nil {
        l1Val := math.MaxInt64
        if l1 != nil {l1Val = l1.Val}
        l2Val := math.MaxInt64
        if l2 != nil {l2Val = l2.Val}
        if l1Val < l2Val {
            tail.Next = &ListNode{Val: l1Val}
            l1 = l1.Next
        } else {
            tail.Next = &ListNode{Val: l2Val}
            l2 = l2.Next
        }
        tail = tail.Next
    }
    return head.Next
}