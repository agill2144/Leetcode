/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    res  := math.MinInt64
    var dfs func(r *TreeNode) int
    dfs = func(r *TreeNode) int {
        // base
        if r == nil {return 0}

        // logic
        left := dfs(r.Left)
        right := dfs(r.Right)
        leftRoot := left+r.Val
        rightRoot := right+r.Val
        all := left+right+r.Val
        res = max(res, max(leftRoot,max(rightRoot,max(r.Val, all))))
        return max(leftRoot,max(rightRoot, r.Val))
    }
    dfs(root)
    return res
}