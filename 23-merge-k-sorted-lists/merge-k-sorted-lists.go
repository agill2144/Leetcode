/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// approach: divide and conquer / merge sort ( recursive )
// tc = o(N * logk)
// sc = o(k) for recursive depth
func mergeKLists(lists []*ListNode) *ListNode {
    k := len(lists)
    if k == 0 {return nil}
    var dfs func(left, right int) *ListNode
    dfs = func(left, right int) *ListNode {
        // base
        if left == right {return lists[left]}
        if left > right {return nil}

        // logic
        mid := left + (right-left)/2
        l1 := dfs(left, mid)
        l2 := dfs(mid+1, right)
        return merge2Lists(l1, l2)
    }
    return dfs(0, k-1)
}





// approach: brute force
// merge lists from 1 to n-1 into lists[0]
// return lists[0]
// N = n*k ( i.e total number of nodes in the final list )
// tc = k*N
// sc = o(1)
// func mergeKLists(lists []*ListNode) *ListNode {
//     k := len(lists)
//     if k == 0 {return nil}
//     head := lists[0]
//     for i := 1; i < len(lists); i++ { // k
//         head = merge2Lists(head, lists[i]) // n
//     }
//     return head
// }

func merge2Lists(l1, l2 *ListNode) *ListNode {
    head := &ListNode{Val: 0}
    tail := head
    for l1 != nil || l2 != nil {
        l1Val := math.MaxInt64
        if l1 != nil {l1Val = l1.Val}
        l2Val := math.MaxInt64
        if l2 != nil {l2Val = l2.Val}

        if l1Val == math.MaxInt64 && l2Val == math.MaxInt64 {break}

        if l1Val < l2Val {
            tail.Next = l1
            tail = tail.Next
            l1 = l1.Next
        } else {
            tail.Next = l2
            tail = tail.Next
            l2 = l2.Next
        }
    }
    return head.Next
}