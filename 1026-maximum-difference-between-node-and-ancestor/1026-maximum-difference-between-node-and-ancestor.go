/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
    approach: top down using postorder ( maintain min and max for each root to leaf path )
    - We just need to find a min and max for each root->leaf path
    - once we reached a nil node from leaf, we can calculate the diff between max-min and update a maxDiff
    time := o(n)
    space: o(n) : recursion stack, skewed tree worse case has n nodes, therefore n calls paused in recursion stack at any given time
*/
func maxAncestorDiff(root *TreeNode) int {
    maxDiff := 0
    var dfs func(cMax, cMin int, r *TreeNode) 
    dfs = func(cMax, cMin int, r *TreeNode) {
        // base
        if r == nil {
            return
        }
        
        // logic
        cMax = max(cMax, r.Val)
        cMin = min(cMin, r.Val)
        dfs(cMax, cMin, r.Left)
        dfs(cMax, cMin, r.Right)
        maxDiff = max(maxDiff, cMax-cMin)

    }
    dfs(math.MinInt64, math.MaxInt64, root)
    return maxDiff
}


func max(x, y int) int {
    if x > y {return x}
    return y
}
func min(x, y int) int {
    if x < y {return x}
    return y
}