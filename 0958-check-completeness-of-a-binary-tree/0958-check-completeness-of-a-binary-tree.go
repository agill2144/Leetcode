/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCompleteTree(root *TreeNode) bool {
    if root == nil {return true}

    nullSeen := false
    q := []*TreeNode{root}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        if dq != nil && nullSeen {return false}
        if dq == nil {nullSeen = true; continue}
        
        q = append(q, dq.Left)
        q = append(q, dq.Right)
    }
    return true
}