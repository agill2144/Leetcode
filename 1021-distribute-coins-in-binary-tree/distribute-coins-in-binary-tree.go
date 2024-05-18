/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distributeCoins(root *TreeNode) int {
    moves := 0
    var dfs func (r *TreeNode) int
    dfs = func(r *TreeNode) int {
        // base
        if r == nil {return 0}

        // logic
        left := dfs(r.Left)
        right := dfs(r.Right)
        extra := left+right+r.Val-1
        moves += abs(left)+abs(right)+r.Val-1        
        return extra
    }
    dfs(root)
    return moves
}


func abs(x int) int {
    if x < 0 {return x*-1}
    return x
}