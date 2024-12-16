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
	if root == nil {return root}
    r := root
    for r != nil {
        curr := r
        for curr != nil {
            if curr.Left != nil && curr.Right != nil {
                curr.Left.Next = curr.Right
            }
            if curr.Right != nil && curr.Next != nil {
                if curr.Next.Left != nil {
                    curr.Right.Next = curr.Next.Left
                } else if curr.Next.Right != nil {
                    curr.Right.Next = curr.Next.Right
                }
            }
            curr = curr.Next
        }
        r = r.Left
    }
    return root
}