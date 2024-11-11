/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    haSize := 0
    curr := headA
    for curr != nil {haSize++ ; curr = curr.Next}
    hbSize := 0
    curr = headB
    for curr != nil {hbSize++ ; curr = curr.Next}
    for haSize > hbSize {headA = headA.Next; haSize--}    
    for hbSize > haSize {headB = headB.Next; hbSize--}    

    for headA != nil && headB != nil{
        if headA == headB {return headA}
        headA = headA.Next
        headB = headB.Next
    }
    return nil
}