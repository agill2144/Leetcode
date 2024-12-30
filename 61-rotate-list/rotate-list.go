/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil || k < 1 {return head}
    n := 0
    curr := head
    for curr != nil {
        n++
        curr = curr.Next
    }
    if k % n == 0 {return head}    
    // n = 5
    // k = 12
    // then 1st rotation would rotate all 5 nodes back into same place
    // now num of rotations left = 12-5 = 7
    // then 2nd rotation would rotate all 5 nodes back into same place
    // now num of rotations left = 7-5 = 2
    // therefore we only need to rotate 2 nodes 
    // which is remainder of k%n
    // "how many more rotations are left, after we have done full n size rotation"
    // if k was smaller?
    // n = 5
    // k = 2
    // 2%5 = 2, same as original k
    k %= n


    // reverse entire LL
    head, _ = reverse(head)


    // reverse first k nodes
    curr = head
    size := 1
    for curr != nil && size != k {
        curr = curr.Next
        size++
    }
    head2 := curr.Next
    curr.Next = nil
    head, tail := reverse(head)

    // reverse the 2nd set of nodes
    head2, _ = reverse(head2)
    tail.Next = head2
    return head
}


func reverse(head *ListNode) (*ListNode, *ListNode) {
    if head == nil || head.Next == nil {return head, head}
    var prev *ListNode
    curr := head
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev, head
}