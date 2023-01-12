/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
    freqMap := map[int]int{0:1}
    count := 0
    
    var dfs func(r *TreeNode, sum int)
    dfs = func(r *TreeNode, sum int) {
        // base 
        if r == nil {return}
        // logic
        sum += r.Val
        diff := sum-targetSum
        if c , ok := freqMap[diff]; ok && c > 0 {count+=c}
        freqMap[sum]++
        
        
        dfs(r.Left, sum)
        dfs(r.Right, sum)
        freqMap[sum]--
    }
    dfs(root, 0)
    return count
}