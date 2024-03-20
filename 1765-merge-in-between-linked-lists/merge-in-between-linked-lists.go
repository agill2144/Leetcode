/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    var (
        prev *ListNode = nil
        athNode *ListNode
        bthNode *ListNode
    )
    curr := list1
    // its not count from 1, its count from 0 :face_palm
    idx := 0
    for curr != nil && (athNode == nil || bthNode == nil) {
        if idx == a {
            athNode = prev
        }
        if idx == b {
            bthNode = curr.Next
        }
        idx++
        prev = curr
        curr = curr.Next
    }
    fmt.Println(athNode, bthNode)
    tailOfList2 := list2
    for tailOfList2.Next != nil {tailOfList2 = tailOfList2.Next}
    athNode.Next = list2
    tailOfList2.Next = bthNode
    return list1
}