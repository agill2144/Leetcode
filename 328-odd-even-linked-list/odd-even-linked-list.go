/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
    - split the even LL into a new dummy chain
    - in the main loop; focus on grouping odd nodes into curr ptr
    - while the even nodes will be appended to dummy head
    - at the end of the loop we will have splitted the LL into 2
    - head = has collected all the odd index position nodes
    - dummy = has colleced all the even index position nodes
    - our initial curr ptr will be sitting at the tail of orignal LL (ones with all odd idx position nodes)
    - simply connect the 2 linkedlists again!
    - oddTail.next(curr ptr; odd nodes) = dummy.next( even nodes )
    time = o(n)
    space = o(1); dummyNode is constant; we reused all the nodes, created no new nodes
*/
func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {return head}
    curr := head
    dummy := &ListNode{Val: 0}
    eTail := dummy
    for curr != nil && curr.Next != nil {
        next := curr.Next
        nextNext := next.Next

        eTail.Next = next
        curr.Next = nil
        eTail = eTail.Next
        eTail.Next = nil

        curr.Next = nextNext
        if nextNext != nil {
            curr = nextNext
        }
    }
    curr.Next = dummy.Next
    return head
}