/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countUnivalSubtrees(root *TreeNode) int {
    count := 0
    var dfs func(r *TreeNode) (int, bool)
    dfs = func(r *TreeNode) (int,bool) {
        // base
        if r == nil {return math.MinInt64, true}
        // logic
        left, leftValid := dfs(r.Left)
        right, rightValid := dfs(r.Right)
        if !leftValid || !rightValid {
            return r.Val, false
        }
    
        if left == math.MinInt64 && right == math.MinInt64 {
            count++
            return r.Val, true
        } else if left != math.MinInt64 && right != math.MinInt64 && left == right && right == r.Val{
            count++
            return r.Val, true
        } else if left == math.MinInt64 && right != math.MinInt64 && right == r.Val {
            count++
            return r.Val, true
        } else if right == math.MinInt64 && left != math.MinInt64 && left == r.Val  {
            count++
            return r.Val, true
        }
        return r.Val, false
    }
    dfs(root)
    return count
}