/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*

approach: bottom up dfs
- In a subtree, from a particular root node of that subtree
- if we have a left child
- then that left child becomes the new root
- the r.left.right = current root node
- the r.left.left = current root right node
- then to offically make the left child of this root node the new root node
- we have to disconnect the the current root left and right pointers
- but if we do this top down, then we will lose access to go to next left child and recurse..
- therefore do the above in bottom up inorder fashion
time: o(n) -- worse case in a skewed tree
space: o(h) or o(n) in worse case if we are given a one-sided skewed tree
*/
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
    if root == nil || root.Left == nil {return root}
    var dfs func(r *TreeNode)
    var newRoot *TreeNode
    
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}
        
        // logic
        dfs(r.Left)
        if r.Left != nil {
            r.Left.Right = r
            r.Left.Left = r.Right
            if newRoot == nil {newRoot = r.Left}
            r.Left = nil; r.Right = nil
        }
        dfs(r.Right)
    }
    dfs(root)
    return newRoot
}