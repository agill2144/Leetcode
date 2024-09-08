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
    out := [][]int{}
    even := []*TreeNode{}
    odd := []*TreeNode{root}
    for len(odd) != 0 || len(even) != 0 {
        // process odd first
        level := []int{}
        for len(odd) != 0 {
            top := odd[len(odd)-1]
            odd = odd[:len(odd)-1]
            level = append(level, top.Val)
            if top.Left != nil {even = append(even, top.Left)}
            if top.Right != nil {even = append(even, top.Right)}
        }
        out = append(out, level)
        level = []int{}
        for len(even) != 0 {
            top := even[len(even)-1]
            even = even[:len(even)-1]
            level = append(level, top.Val)
            if top.Right != nil {odd = append(odd, top.Right)}
            if top.Left != nil {odd = append(odd, top.Left)}
        }
        if len(level) > 0 {
            out = append(out, level)
        }
    }
    return out
}