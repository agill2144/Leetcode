/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    var dfs func(r *TreeNode) (int, bool)
    dfs = func(r *TreeNode) (int, bool) {
        // base
        if r == nil {return 0, true}

        // logic
        left, lok := dfs(r.Left)
        right, rok := dfs(r.Right)
        if !lok || !rok {return -1, false}
        if abs(right-left) > 1 {return -1, false}
        return max(left, right) + 1, true
    }
    _, ok := dfs(root)
    return ok
}

func abs(x int) int {if x < 0 {return x*-1} else {return x}}