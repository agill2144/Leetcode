/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findLeaves(root *TreeNode) [][]int {
    result := [][]int{}
    var dfs func(r *TreeNode) int
    dfs = func(r *TreeNode) int {
        // base
        if r == nil {return 0}
        
        // logic
        left := dfs(r.Left)
        right := dfs(r.Right)
        maxHeight := max(left, right)
        if len(result) == maxHeight {
            result = append(result, []int{})
        }
        result[maxHeight] = append(result[maxHeight],r.Val)
        return maxHeight+1
    }
    dfs(root)
    return result
}

func max(x, y int) int {
    if x > y {return x}
    return y
}