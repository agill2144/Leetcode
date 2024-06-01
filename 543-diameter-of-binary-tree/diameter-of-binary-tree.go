/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    var dfs func(r *TreeNode) int
    maxDepth := 0
    dfs = func(r *TreeNode) int {
        // base
        if r == nil {return 0}

        // logic
        left := dfs(r.Left)
        right := dfs(r.Right)
        maxDepth = max(maxDepth, left+right)
        return max(left,right)+1

    }
    dfs(root)
    return maxDepth
}