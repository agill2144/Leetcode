/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    st := []*TreeNode{}
    out := []int{}
    for root != nil || len(st) != 0 {
        for root != nil {
            st = append(st, root)
            root = root.Left
        }
        root = st[len(st)-1]
        st = st[:len(st)-1]
        out = append(out, root.Val)
        root = root.Right
    }
    return out
}