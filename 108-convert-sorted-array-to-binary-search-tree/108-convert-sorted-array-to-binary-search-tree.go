/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    var dfs func(left, right int) *TreeNode
    dfs = func(left, right int) *TreeNode{
        // base
        if left > right {return nil}
        
        // logic
        mid := left + (right-left)/2
        root := &TreeNode{Val: nums[mid]}
        root.Left = dfs(left, mid-1)
        root.Right = dfs(mid+1, right)
        return root
    }
    return dfs(0, len(nums)-1)
}