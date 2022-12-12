/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    // 1->2->3->4
    // insert(5)
    
    mergedHead := &ListNode{Val: 0}
    tail := mergedHead
    // 0
    
    l1 := list1
    l2 := list2
    
    for l1 != nil || l2 != nil {
        var l1Val int = 101
        if l1 != nil {
            l1Val = l1.Val
        }
        
        var l2Val int = 101
        if l2 != nil {
            l2Val = l2.Val
        }
        
        if l1Val != 101 && l2Val != 101 {
            if l1Val < l2Val {
                tail.Next = &ListNode{Val: l1Val}
                tail = tail.Next
                l1 = l1.Next
            } else {
                tail.Next = &ListNode{Val: l2Val}
                tail = tail.Next
                l2 = l2.Next
            }
        } else if l1Val != 101 && l2Val == 101 {
            tail.Next = &ListNode{Val: l1Val}
            tail = tail.Next
            l1 = l1.Next
        } else {
            tail.Next = &ListNode{Val: l2Val}
            tail = tail.Next
            l2 = l2.Next
        }
    }
    
    return mergedHead.Next
}