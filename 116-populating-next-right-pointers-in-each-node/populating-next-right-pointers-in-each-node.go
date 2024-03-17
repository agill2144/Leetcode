/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

/*
    approach: level order using BFS
    - connect current dq'd node to next node in queue
        - only IF NEXT NODE IS IN THE SAME LEVEL
        - i.e do not connect to next node in queue if current dq'd node is the last node in level
    
    timn = o(n)
    space = o(maxWidthOfTree); o(n/2); o(n)
*/
func connect(root *Node) *Node {
    if root == nil {return nil}
	q := []*Node{root}
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if qSize > 1 && len(q) != 0 {
                dq.Next = q[0]
            }
            if dq.Left != nil {
                q = append(q, dq.Left)
            }
            if dq.Right != nil {
                q = append(q, dq.Right)
            }
            qSize--
        }
    }
    return root
}