/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distributeCoins(root *TreeNode) int {
    // 1 coin  == 1 move
    // 2 coins == 2 moves
    // calc extra coins after leaving a coin at current node
    // thats the number of moves we need to make
    // because we need to move $extra coins
    // and $extra coins == $extra moves
    moves := 0
    var dfs func (r *TreeNode) int
    dfs = func(r *TreeNode) int {
        // base
        if r == nil {return 0}

        // logic
        left := dfs(r.Left)
        right := dfs(r.Right)
        extras := left+right+r.Val-1
        moves += abs(left)+abs(right)+r.Val-1
        return extras
    }
    dfs(root)
    return moves
}


func abs(x int) int {
    if x < 0 {return x*-1}
    return x
}