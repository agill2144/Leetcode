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
        if left > right {return nil}
        if left == right {return lists[left]}

        // logic
        mid := left + (right-left)/2
        l1 := dfs(left, mid)
        l2 := dfs(mid+1, right)
        return merge2Lists(l1, l2)
    }
    return dfs(0, len(lists)-1)
}

func merge2Lists(l1,l2 *ListNode) *ListNode {
    head := &ListNode{Val:0}
    tail := head
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            tail.Next = l1
            l1 = l1.Next
        } else {
            tail.Next = l2
            l2 = l2.Next
        }
        tail = tail.Next
    }
    for l1 != nil {
        tail.Next = l1
        l1 = l1.Next
        tail = tail.Next
    }
    for l2 != nil {
        tail.Next = l2
        l2 = l2.Next
        tail = tail.Next
    }
    return head.Next
}