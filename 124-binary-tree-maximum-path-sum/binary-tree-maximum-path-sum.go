/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
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
        // 4 possible sums
        //1. curr node + leftSum + rightSum
        leftRightAndCurr := left+right+r.Val
        //2. curr node + leftSum
        leftAndCurr := left+r.Val
        //3. curr node + rightSum
        rightAndCurr := right+r.Val
        // 4. curr node on its own
        // take the max sum out of the above 4 possibilites
        maxAtThisNode := max(leftRightAndCurr, max(leftAndCurr, max(rightAndCurr, r.Val)))
        // save the maxSum 
        maxSum = max(maxSum, maxAtThisNode)
        // return a valid path back to parent, 
        // a valid path going back to parent will never contain all nodes ( curr+left+right )
        // it will either be curr node || curr+left || curr+right
        return max(r.Val, max(leftAndCurr,rightAndCurr))
    }
    dfs(root)
    return maxSum
}