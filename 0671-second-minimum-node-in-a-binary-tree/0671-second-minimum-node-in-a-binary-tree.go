/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findSecondMinimumValue(root *TreeNode) int {
    min := math.MaxInt64
    secondMin := math.MaxInt64
    
    var dfs func(r *TreeNode) 
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}
        
        // logic
        if r.Val < min {
            secondMin = min
            min = r.Val
        } 
        if r.Val < secondMin && r.Val > min {
            secondMin = r.Val
        }
        dfs(r.Left)
        dfs(r.Right)
    }
    dfs(root)
    if secondMin == math.MaxInt64 { return -1 }
    return secondMin
}