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
        if curr.Val <= curr.Next.Val && x >= curr.Val && x <= curr.Next.Val {
            break
        } else if curr.Val > curr.Next.Val {
            if x >= curr.Val || x <= curr.Next.Val {
                break
            }
        }
        curr = curr.Next
    }
    next := curr.Next
    curr.Next = newNode
    newNode.Next = next
    return aNode
}