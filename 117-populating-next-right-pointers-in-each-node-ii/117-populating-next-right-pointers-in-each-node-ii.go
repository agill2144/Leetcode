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
   if root == nil {
		return nil
	}
	q := []*Node{root}
	for len(q) != 0 {
		qSize := len(q)
		// [2-3]
		// qSize = 2
		for i := 0 ; i < qSize; i++{
			dq := q[0]
			q = q[1:]
			if len(q) > 0 && i+1 != qSize {
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