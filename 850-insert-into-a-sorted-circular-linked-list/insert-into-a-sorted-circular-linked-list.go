/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 * }
 */

/*
    approach: linear scan
    - entry node = node given in input
    - start node = smallest node in the LL, this is where sorting actually starts from
    - end node = largest node in the LL
    - there are 3 places we would want to insert x at
        1. in between 2 sorted nodes
        - meaning curr <= next
        - and x >= curr and x <= next
        - therefore x can be placed in-between curr and next

        2. after the ending node OR before the starting node
        ** IMPORTANT **
        - how do do we detect we are at the ending / starting node?
        - ending node is supposed to be the largest
        - if curr node value > next node value
        - it means, we are not looking at 2 sorted nodes
        - and the only time thats possible is when at the end of a sorted list
        - end item will be > than start item in a sorted list, which means
        - this curr node is the ending node / ending item in a sorted list
        - and next node is the starting node / starting item in a sorted list
        - we can insert x after ending node if x >= ending node
        - OR we can insert x before starting node if x <= starting node
        - therefore when we are at a ending node
        - we can insert x IF
            - x >= ending node OR x <= starting node
        - but could x be smaller than ending and greater than starting node ?
            - ending = 5
            - starting = 1
            - entry = 3
            - x = 2
            - we need to insert x before entry node

        3. before entry node
        - could x not be greater than ending node and also not be smaller than starting node?
        - yes!
          ---<--<--<--
         |           |
        - 3->4->5->1-
        - x = 2
        - entry node = 3; 
        - ending node= 5; 
        - starting node = 1
        - 2 does not fit between 3->4, 4->5 or 5->1(ending->starting)
        - 2 fits BEFORE entry node ( i.e after starting node )

    - we need to traverse the LL, ofc
    - when do we stop traversing? its circular, whats our exit case?
    - we can stop when we lapped the entire LL once
    - meaning curr.Next == entry node
    
    - 1. is easy to check and handle
    - if curr and next are sorted, we can check if x can sit in between
    - if yes, we can break out, and use curr ptr to insert newNode and connect newNode to next node

    - 2. how do we know we are at ending / starting node?
    - when 2 nodes are not sorted
    - meaning if its not 1, then this is defintely the case
    - we can then check, can x come after ending node or can x be before starting node?
    - if x >= curr or x <= next, if yes, break and we can insert new node after curr ptr for both cases

    - 3. before entry node
    - this automatically happens when our while loop terminates
    - meaning, we could not insert in the middle, or after ending node or before starting node
    - therefore we need to insert before ENTRY node
    - when while loop breaks, our curr ptr is sitting at starting node, and next node is entry node
    - we want to insert before entry node
    - meaning after curr ptr

    - therefore all 3 cases can be handled by breaking and inserting after curr ptr
    - wherever curr ptr ends

    tc = o(n)
    sc = o(1)
*/

func insert(aNode *Node, x int) *Node {
    newNode := &Node{Val: x}
    if aNode == nil {newNode.Next = newNode; return newNode}
    curr := aNode
    for curr.Next != aNode{
        next := curr.Next
        if curr.Val <= next.Val {
            if x >= curr.Val && x <= next.Val {break}
        } else {
            // curr = ending node
            // next = starting node
            if x >= curr.Val || x <= next.Val {break}
        }
        curr = next
    }
    next := curr.Next
    curr.Next = newNode
    newNode.Next = next
    return aNode
}