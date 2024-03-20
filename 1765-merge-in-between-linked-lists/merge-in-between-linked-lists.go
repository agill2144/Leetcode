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
        next := curr.Next
        // we are removing nodes from idx a to b ( including ath idx and bth idx)
        // thats why we need a's prev, and b's next
        if idx == a { athNode = prev }
        if idx == b { bthNode = next }
        idx++
        prev = curr
        curr = curr.Next
    }
    tail := list2
    for tail.Next != nil {tail = tail.Next}
    athNode.Next = list2
    tail.Next = bthNode
    return list1
}