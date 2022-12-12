/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// in = left - root - right
func inorderTraversal(root *TreeNode) []int {
    out := []int{}
    stack := []*TreeNode{}
    for root != nil || len(stack) != 0 {
        // left
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        // process root
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        out = append(out, root.Val)
        
        // right
        root = root.Right
    }
    return out
}