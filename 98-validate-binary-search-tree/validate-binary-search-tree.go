/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 /*
    inorder traversal will traverse in a sorted order ( if the bst was sorted )
    whenever we see a bst , think sorted == inorder
    bst == sorted == inorder

    approach: inorder traversal with a gloabl prev
    - running a inorder traversal will give us a sorted order
    - if we were checking whether a array is sorted or not, we will each element will previous element
    - therefore we will maintain a prev ptr that tells us the previous value
    - 
 */
func isValidBST(root *TreeNode) bool {
    var dfs func(r *TreeNode, min, max int) bool
    dfs = func(r *TreeNode, min, max int) bool {
        // base
        if r == nil {return true}
        // logic
        if r.Val <= min || r.Val >= max {return false}
        if !dfs(r.Left, min, r.Val) {return false}
        if !dfs(r.Right, r.Val, max) {return false}
        return true
    }
    return dfs(root, math.MinInt64, math.MaxInt64)
}