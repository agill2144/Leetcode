/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    aSize := 0
    bSize := 0
    curr := headA
    for curr != nil {aSize++; curr=curr.Next}
    curr = headB
    for curr != nil {bSize++; curr=curr.Next}
    for aSize > bSize {headA = headA.Next; aSize--}
    for bSize > aSize {headB = headB.Next; bSize--}
    for headA != nil && headB != nil {
        if headA == headB  {return headA}
        headA = headA.Next
        headB = headB.Next
    }
    return nil
}