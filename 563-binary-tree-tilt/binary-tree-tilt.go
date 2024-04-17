/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
    tilt := 0
    var dfs func(r *TreeNode) int
    dfs = func(r *TreeNode) int {
        // base
        if r == nil {return 0}

        // logic
        left := dfs(r.Left)
        right := dfs(r.Right)
        diff := abs(left-right)
        tilt += diff
        return left+right+r.Val
    }
    dfs(root)
    return tilt
}

func abs(x int) int {
    if x < 0 {return x * -1}
    return x
}