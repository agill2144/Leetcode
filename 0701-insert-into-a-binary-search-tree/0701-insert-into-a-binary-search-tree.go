/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{Val: val}
    }
    
    r := root
    for r != nil {
        if val > r.Val {
            if r.Right != nil {
                r = r.Right
            } else {
                r.Right = &TreeNode{Val: val}
                break
            }
        } else {
            if r.Left != nil {
                r = r.Left
            } else {
                r.Left = &TreeNode{Val: val}
                break
            }
        }
    }
    return root
}