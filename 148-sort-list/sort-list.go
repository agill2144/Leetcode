/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    var dfs func(h *ListNode) *ListNode
    dfs = func(h *ListNode) *ListNode {
        // base
        if h == nil || h.Next == nil {return h}

        // logic
        slow := h
        fast := h
        var prev *ListNode
        for fast != nil && fast.Next != nil {
            prev = slow
            slow = slow.Next
            fast = fast.Next.Next
        }
        prev.Next = nil
        h1 := dfs(h)
        h2 := dfs(slow)
        dummy := &ListNode{Val:-1}
        tail := dummy 
        l1, l2 := h1,h2
        for l1 != nil && l2 != nil {
            l1Next := l1.Next
            l2Next := l2.Next
            if l1.Val < l2.Val {
                l1.Next = nil
                tail.Next = l1
                tail = tail.Next
                l1 = l1Next
            } else {
                l2.Next = nil
                tail.Next = l2
                tail = tail.Next
                l2 = l2Next
            }
        }
        for l1 != nil {
            l1Next := l1.Next
            l1.Next = nil
            tail.Next = l1
            tail = tail.Next
            l1 = l1Next
        }
        for l2 != nil {
            l2Next := l2.Next
            l2.Next = nil
            tail.Next = l2
            tail = tail.Next
            l2 = l2Next
        }
        newHead := dummy.Next
        return newHead
    }
    return dfs(head)
}