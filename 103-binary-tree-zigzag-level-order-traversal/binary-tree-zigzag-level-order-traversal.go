/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil {return nil}
    st1 := []*TreeNode{root} // left->right
    st2 := []*TreeNode{} // right->left
    out := [][]int{}    
    for len(st1) != 0 || len(st2) != 0 {
        level := []int{}
        for len(st1) != 0 {
            top := st1[len(st1)-1]
            st1 = st1[:len(st1)-1]
            level = append(level, top.Val)
            if top.Left != nil {
                st2 = append(st2, top.Left)
            }
            if top.Right != nil {
                st2 = append(st2, top.Right)
            }
        }
        out = append(out, level)
        level = []int{}
        for len(st2) != 0 {
            top := st2[len(st2)-1]
            st2 = st2[:len(st2)-1]
            level = append(level, top.Val)
            if top.Right != nil {
                st1 = append(st1, top.Right)
            }
            if top.Left != nil {
                st1 = append(st1, top.Left)
            }
        }
        if len(level) > 0 {
            out = append(out, level)
        }
    }
    return out
}