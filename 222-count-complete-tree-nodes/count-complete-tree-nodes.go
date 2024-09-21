/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// brute force
func countNodes(root *TreeNode) int {
    count := 0
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode){
        // base
        if r == nil {return}

        //logic
        count++
        dfs(r.Left)
        dfs(r.Right)
    }
    dfs(root)
    return count
}