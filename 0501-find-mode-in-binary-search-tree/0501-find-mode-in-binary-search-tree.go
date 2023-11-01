/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// collect all dupes 
func findMode(root *TreeNode) []int {
    freqMap := map[int]int{}
    maxCount := 0
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}
        
        // logic
        dfs(r.Left)
        freqMap[r.Val]++
        if freqMap[r.Val] > maxCount {maxCount = freqMap[r.Val]}
        dfs(r.Right)
    }
    dfs(root)
    out := []int{}
    for k, v := range freqMap {
        if v == maxCount {out = append(out, k)}
    }
    return out
}