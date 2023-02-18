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
    q := []*Node{root}
    for len(q) != 0 {
        qSize := len(q)
        for i := 0; i < qSize; i++ {
            dq := q[0]
            q = q[1:]
            if i != qSize-1 {
                dq.Next = q[0]
            }
            if dq.Left != nil {
                q = append(q, dq.Left)
            }
            if dq.Right != nil {
                q = append(q, dq.Right)
            }
        }
    }
    return root
}