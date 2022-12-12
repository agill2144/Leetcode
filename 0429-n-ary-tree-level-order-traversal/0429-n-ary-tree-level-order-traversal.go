/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
    if root == nil {return nil}
    q := []*Node{root}
    out := [][]int{}
    for len(q) != 0 {
        qSize := len(q)
        level := []int{}
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            level = append(level, dq.Val)
            if dq.Children != nil && len(dq.Children) > 0 {
                q = append(q, dq.Children...)
            }
            qSize--
        }
        out = append(out, level)
    }
    return out

}