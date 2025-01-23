/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    idx := 0
    var dfs func(r *TreeNode) (*TreeNode)
    dfs = func(r *TreeNode) (*TreeNode) {
        // base
        if r == nil {return nil}


        // logic
        left := dfs(r.Left)
        if left != nil {return left}
        if idx == k-1 {return r}
        idx++
        return dfs(r.Right)
    }
    return dfs(root).Val
    
}