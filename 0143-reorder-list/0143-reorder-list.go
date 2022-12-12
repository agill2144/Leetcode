/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 /*
    approach 1: double ended queue
    
    Loop over LL, while saving each node in a double ended queue
    Then make the head nil ( reset )
    

    1. Initialize head with a dummy node and use a current ptr to act as the tail while we are appending
    while we have items in double ended queue:
        1. current.Next = Pop from front

        if we still have items:
            2. current.Next = Pop from back ( double check the dq still has elements remaining because the first operation will run out )
        
    time: o(n)
    space: o(n) because of the extra double ended queue
        
        
        
    approach 2: 
    - split LL into 2 halves
    - reverse second half of the LL ( in place )
    - Then we will have 2 LL's ( start and second-reversed )
    - Merge 2 LL's starting with first and then second
    
    time: o(n)
    space: o(1)
*/
func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil {
        return 
    }
    slow := head
    fast := head
    var prev *ListNode
    for fast != nil && fast.Next != nil {
        prev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    prev.Next = nil
    slow = reverse(slow)
    // fmt.Println(slow.Val)
    p1 := head
    p2 := slow
    p1Tail := p1 // when we are out of nodes in p1, we will use this tail ptr to append new nodes from p2
    for p1 != nil || p2 != nil {
        var p1Next *ListNode
        if p1 != nil {
            p1Next = p1.Next
        }
        var p2Next *ListNode
        if p2 != nil {
            p2Next = p2.Next
        }
        if p1 != nil {
            p1.Next = p2
            p2.Next = p1Next
        } else if p2 != nil {
            p1Tail.Next = p2
        }
        p1Tail = p2
        p1 = p1Next
        p2 = p2Next
    }
    
}

func reverse(head *ListNode) *ListNode{
    if head == nil || head.Next == nil {
        return head
    }
    var prev *ListNode
    curr := head
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}