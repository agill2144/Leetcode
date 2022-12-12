/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
    approach: bottom up : inverse of max depth
    time: o(n)
    space: o(n) ; worse case we have a skewed tree
*/
func minDepth(root *TreeNode) int {
    if root == nil {return 0}
    var dfs func(r *TreeNode) (int, bool)
    dfs = func(r *TreeNode) (int, bool ){
        // base
        if r == nil {return math.MaxInt64, false}
        
        // logic
        left, leftFound := dfs(r.Left)
        right, rightFound := dfs(r.Right)
        
        if !leftFound && !rightFound {
            return 1, true // leaf node
        } else if leftFound && rightFound {
            return min(left, right)+1, true
        } else if leftFound {
            return left+1, true
        } else {
            return right+1, true
        }
    }
    m, _ := dfs(root)
    return m
}

func min(x, y int) int{
    if x < y {return x}
    return y
}
