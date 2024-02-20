/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
    approah: level order using bfs
    - level order using

    time = o(n)
    space = o(maxWidth of tree) ; i.e last level ; i.e o(n/2); i.e o(n)
*/
func levelOrder(root *TreeNode) [][]int {
    if root == nil {return nil}
    out := [][]int{}
    q := []*TreeNode{root}
    for len(q) != 0 {
        qSize := len(q)
        level := []int{}
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            level = append(level, dq.Val)
            if dq.Left != nil {
                q = append(q, dq.Left)
            }
            if dq.Right != nil {
                q = append(q, dq.Right)
            }
            qSize--
        }
        out = append(out,level)
    }
    return out
}