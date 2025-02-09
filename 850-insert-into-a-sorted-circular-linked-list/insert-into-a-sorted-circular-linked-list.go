/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 * }
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
            if x >= curr.Val && x <= next.Val {
                break
            }
        } else if curr.Val > next.Val {
            if x >= curr.Val || x <= next.Val {
                break
            }
        }
        curr = next
    }
    next := curr.Next
    curr.Next = newNode
    newNode.Next= next
    return aNode
}