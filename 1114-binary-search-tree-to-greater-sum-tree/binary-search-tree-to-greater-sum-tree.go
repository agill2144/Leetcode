/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // approach: do inorder backwards
 // inorder on a bst gives elements in a asc order
 // inorder backwards on a bst will give us elements in a desc order
func bstToGst(root *TreeNode) *TreeNode {
    var dfs func(r *TreeNode) 
    rSum := 0
    dfs = func(r *TreeNode)  {
        // base
        if r == nil {return }

        // logic
        dfs(r.Right)
        r.Val += rSum
        rSum = r.Val
        dfs(r.Left)
    }
    dfs(root)
    return root
}