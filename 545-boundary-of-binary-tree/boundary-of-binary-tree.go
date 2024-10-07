/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func boundaryOfBinaryTree(root *TreeNode) []int {
    if root == nil {return nil}
    isLeaf := func(r *TreeNode) bool {
        return r.Left == nil && r.Right == nil
    }
    out := []int{root.Val}
    leftView := root.Left
    for leftView != nil && !isLeaf(leftView) {
        out = append(out, leftView.Val)
        if leftView.Left != nil {leftView =leftView.Left} else {leftView = leftView.Right}
    }
    // collect leaves
    var leaves func(r *TreeNode)
    leaves = func(r *TreeNode) {
        // base
        if r == nil {return}
        
        // logic
        if isLeaf(r) && r != root {
            out = append(out, r.Val)
        }
        leaves(r.Left)
        leaves(r.Right)
    }
    leaves(root)

    // collect right
    right := []int{}
    rightView := root.Right
    for rightView != nil && !isLeaf(rightView) {
        right = append(right, rightView.Val)
        if rightView.Right != nil {rightView = rightView.Right } else {rightView = rightView.Left}
    }
    for i := len(right)-1; i >= 0; i-- {
        out = append(out, right[i])
    }
    return out
}