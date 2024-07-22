/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
    - using 2 ptrs
    - create a dist between 2 ptrs such that they are n dist away from each other
    - p1 = head
    - p2 = head
    - n = 2
    - move p2 2(n) distance away from p1
    - therefore the dist between p1 and p2 will be exactly n
    - now in another loop; move p1 and p2 until p2 reaches nil
    - when p2 reaches nil, p1 will n steps behind, thereby giving us access to nth node from end
    - since this is not a DLL, maintain a prev ptr behing p1, so that we can remove this(p1 or nth) node
    time = o(n)
    space = o(1)
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil {return head}
    p1 := head
    p2 := head
    
    // create a distance between p1 and p2 of size n
    for i := 0; i < n; i++ {
        p2 = p2.Next
    }

    // now move p1 and p2 until p2 reaches the end
    // remember that p1 and p2 are n distance away from each other
    // so p2 will reach end first
    // and when it does, p1 will be n steps behind ( since they were n distance away from each other )
    // meaning p1 will at the nth node to delete
    var prev *ListNode
    for p2 != nil {
        prev = p1
        p2 = p2.Next
        p1 = p1.Next
    }
    if prev != nil {
        prev.Next = p1.Next
        p1.Next = nil
    } else {
        // p1 is at head node
        newHead := p1.Next
        head.Next = nil
        head = newHead
    }
    return head
}