/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // bottom up
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
    var dfs func(r *TreeNode) 
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}
        // logic
        dfs(r.Left)
        dfs(r.Right)

        if isLeaf(r.Left) && r.Left.Val == target {
            r.Left = nil
        }

        if isLeaf(r.Right) && r.Right.Val == target {
            r.Right = nil
        }
    }
    dfs(root)
    if root.Val == target && isLeaf(root) {
        return nil
    }
    return root
}

func isLeaf(r *TreeNode) bool {
    if r == nil {return false}
    return r.Left == nil && r.Right == nil
}