/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }


7 7
2 2
11 20
4 24
13 13
1 1
4 5
8 26
5 55

 */
func maxPathSum(root *TreeNode) int {
    maxSum := math.MinInt64
    var dfs func(r *TreeNode) int
    dfs = func(r *TreeNode) int {
        // base
        if r == nil {return 0}

        // logic
        left := dfs(r.Left)
        right := dfs(r.Right)
        
        total := left+right+r.Val
        leftAndCurr := left+r.Val
        rightAndCurr := right+r.Val
        maxAtThisNode := max(total, max(leftAndCurr, max(rightAndCurr, r.Val)))
        maxSum = max(maxSum, maxAtThisNode)
        return max(r.Val, max(r.Val+left,r.Val+right))
    }
    dfs(root)
    return maxSum
}