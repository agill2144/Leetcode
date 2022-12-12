/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
    for root != nil {
        if val == root.Val {
            return root 
        }
        if val < root.Val {
            root = root.Left
        } else {
            root = root.Right
        }
    }
    return nil
}