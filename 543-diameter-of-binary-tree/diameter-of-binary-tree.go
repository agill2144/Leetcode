/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    maxDia := 0
    var dfs func(r *TreeNode) int
    dfs = func(r *TreeNode) int {
        // base
        if r == nil {return 0}
        // logic
        left := dfs(r.Left)
        right := dfs(r.Right)
        if r.Left == nil && r.Right == nil {
            return 1
        }
        maxDia = max(maxDia, max(left,max(right, left+right)))
        return max(left,right)+1
    }
    dfs(root)
    return maxDia
}