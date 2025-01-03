/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    getLen := func(head *ListNode) int {
        n := 0
        curr := head
        for curr != nil {
            n++
            curr = curr.Next
        }
        return n
    }
    aSize := getLen(headA)
    bSize := getLen(headB)
    for aSize > bSize {headA=headA.Next; aSize--}
    for bSize > aSize {headB=headB.Next; bSize--}
    for headA != nil && headB != nil{
        if headA == headB {return headA}
        headA = headA.Next
        headB = headB.Next
    }
    return nil
}