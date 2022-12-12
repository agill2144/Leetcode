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
        if r == nil {
            return 0, true
        }
        
        // logic
        left, leftBalanced := dfs(r.Left)
        if !leftBalanced {return 0,false}
        
        right, rightBalanced := dfs(r.Right)
        if !rightBalanced {return 0,false}
        
        if abs(left-right) > 1 {
            return 0,false
        }
        
        return max(left,right)+1, true
    }
    _, ans := dfs(root)
    return ans
}

func abs(x int) int {
    if x < 0 {return x * -1}
    return x
}

func max(x, y int) int {
    if x > y {return x}
    return y
}