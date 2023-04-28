/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
    if nums == nil {return nil}
    var dfs func(left, right int) *TreeNode
    dfs = func(left, right int) *TreeNode {
        // base
        if left > right {return nil}
        
        
        // logic
        // find the max in this window
        maxVal, idx := nums[left], left
        i := left+1
        for i <= right {
            if nums[i] > maxVal {
                maxVal = nums[i]
                idx = i 
            }
            i++
        }
        root := &TreeNode{Val: maxVal}
        root.Left = dfs(left, idx-1)
        root.Right = dfs(idx+1, right)
        return root
    }
    return dfs(0, len(nums)-1)
}