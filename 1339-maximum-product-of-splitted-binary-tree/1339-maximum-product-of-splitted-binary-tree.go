/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxProduct(root *TreeNode) int {
    totalSum := 0
    var getTotalSum func(r *TreeNode)
    getTotalSum = func(r *TreeNode) {
        // base
        if r == nil {return}
        
        // logic
        totalSum += r.Val
        getTotalSum(r.Left)
        getTotalSum(r.Right)
    }
    getTotalSum(root)
    maxProd := 0
    var dfs func(r *TreeNode) int
    dfs = func(r *TreeNode) int {
        // base
        if r == nil {return 0}
        
        // logic
        left := dfs(r.Left)
        right := dfs(r.Right)
        total := left+right+r.Val
        diff := totalSum - total
        
        maxProd = max(maxProd, diff*total)
        
        return left+right+r.Val
    }
    dfs(root)
    return maxProd % 1000000007
}

func max(x, y int) int {
    if x > y {return x}
    return y
}