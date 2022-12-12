/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
    if root == nil {return nil}
    level := root
    c := root
    
    for level != nil {
        for c != nil {
            if c.Left != nil && c.Right != nil {
                c.Left.Next = c.Right
            }
            if c.Next != nil {
                if c.Right != nil {
                    c.Right.Next = c.Next.Left
                }
            }
            c = c.Next
        }
        level = level.Left
        c = level
    }
    return root
}