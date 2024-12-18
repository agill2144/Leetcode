/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 * }
 */
/*

    approach: scan + with some ptrs
    - input node is not always the head node
    - there are 2 possible places where we can insert x
        1. somewhere in the middle of LL
        2. at the end of LL
            - either after last node 
            - or before next node
            - 5->6->7->3 
            - x = 8 ( after 7 )
            - x = 0 ( before 3)
    
    1. somewhere in the middle of LL
    - we can determine whether we can insert in the middle
    - if curr and next are sorted
        - 1->5
        - curr = 1
        - next = 5
        - they are sorted
    - so if curr and next are sorted
    - and x is >= curr and x <= next
    - we can stop and insert our newNode in-between

    2. at the end of LL
    - when are at the end of LL
    - we wont have sorted order!
    - 5->6->7->3 
    - curr = 7
    - next = 3
    - these are not sorted in asc order
    - meaning this is our 2nd case,
    - that is we are trying to insert after last
    - or we need to insert before first node
    - we need to insert after curr if x >= curr.Val
        - 5->6->7->3 
        - curr = 7
        - next = 3
        - x = 8
        - x >= curr; 8 >= 7, yes insert after current
    - OR we need to insert before next if x <= next.Val
        - 5->6->7->3 
        - curr = 7
        - next = 3
        - x = 1
        - x <= next; 1 <= 3; yes, insert before next
        - i.e insert after current 

    - EDGE CASE
    - its possible we dont break out of our loop due to 2 cond checks
    - 3->4->1
    - x = 2
    - curr = 3; next = 4; sorted; cannot insert in-between
    - curr = 4; next = 1; not sorted; 2 >= 4 nope; 2 <= 1; also no
    - curr = 1; but curr.Next = aNode (3)
    - so our main loop breaks
    - when this happens, it means we want to insert after our curr ptr
    - i.e after LL has ran a circle but node to insert is BEFORE entrypoint
        

*/
func insert(aNode *Node, x int) *Node {
    newNode := &Node{Val: x}
    if aNode == nil {
        newNode.Next = newNode
        return newNode
    }

    curr := aNode
    for curr.Next != aNode {
        next := curr.Next
        if curr.Val <= next.Val {
            if x >= curr.Val && x <= next.Val{break}
        } else {
            if x >= curr.Val || x <= next.Val {break}
        }
        curr = next
    }
    next := curr.Next
    curr.Next = newNode
    newNode.Next = next
    return aNode
}