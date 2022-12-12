/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
    if root  == nil {
        return 0
    }
    runningSum := map[int]int{0:1}
    count := 0
    
    var dfs func(r *TreeNode, sum int)
    dfs = func(r *TreeNode, sum int) {
        // base
        if r == nil {return}
        // logic
        sum += r.Val
        if val, ok := runningSum[sum-targetSum]; ok {
            count += val
        }
        runningSum[sum]++
        dfs(r.Left, sum)
        dfs(r.Right, sum)
        runningSum[sum]--
        if runningSum[sum] == 0 {delete(runningSum, sum)}
    }
    dfs(root, 0)
    return count
}