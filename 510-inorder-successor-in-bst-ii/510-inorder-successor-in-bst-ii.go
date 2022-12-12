/**
 * Definition for Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Parent *Node
 * }
 */
// time: o(n)
// space: o(1)

func inorderSuccessor(node *Node) *Node {
    if node == nil {return nil}
    if node.Right != nil {
        curr := node.Right
        for curr.Left != nil {
            curr = curr.Left
        }
        return curr
    }
    var child *Node
    for node != nil {
        if child != nil {
            if child.Val < node.Val {
                return node
            }
        }
        child = node
        node = node.Parent
    }
    return nil
}