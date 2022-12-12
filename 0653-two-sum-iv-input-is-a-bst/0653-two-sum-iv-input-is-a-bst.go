/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
    sortedList := []int{}
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}
        
        // logic
        dfs(r.Left)
        sortedList = append(sortedList, r.Val)
        dfs(r.Right)
    }
    dfs(root)
    left := 0
    right := len(sortedList)-1
    
    for left < right {
        sum := sortedList[left] + sortedList[right]
        if sum == k {
            return true
        } else if sum > k {
            right--
        } else {
            left++
        }
    }
    return false
}