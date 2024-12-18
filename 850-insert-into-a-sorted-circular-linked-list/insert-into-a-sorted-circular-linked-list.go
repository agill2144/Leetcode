/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 * }
 */

func insert(aNode *Node, x int) *Node {
    node := &Node{Val: x}
    if aNode == nil {
        node.Next = node
        return node
    }
    
    curr := aNode
    for curr.Next != aNode {
        next := curr.Next
        if curr.Val <= next.Val {
            if x >= curr.Val && x <= next.Val {
                break
            }
        } else {
            if x >= curr.Val || x <= next.Val {
                break
            }   
        }

        curr = next
    }
    next := curr.Next
    curr.Next = node
    node.Next = next
    return aNode
}