/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
    count := 0
    var dfs func(r *TreeNode, rSum int, path map[int]int)
    dfs = func(r *TreeNode, rSum int, path map[int]int) {
        // base
        if r == nil {return}
        
        // logic
        rSum += r.Val
        diff := rSum-targetSum
        count += path[diff]
        path[rSum]++
        dfs(r.Left, rSum, path)
        dfs(r.Right, rSum, path)
        path[rSum]--
    }
    dfs(root, 0, map[int]int{0:1})
    return count
}