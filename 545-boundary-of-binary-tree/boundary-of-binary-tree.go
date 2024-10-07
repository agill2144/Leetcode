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
    leftView := func(r *TreeNode) {
        for r != nil && !isLeaf(r) {
            out = append(out, r.Val)
            if r.Left != nil {r =r.Left} else {r = r.Right}
        }
    }
    leftView(root.Left)
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
    rightView := func(r *TreeNode) {
        for r != nil && !isLeaf(r) {
            right = append(right, r.Val)
            if r.Right != nil {r = r.Right } else {r = r.Left}
        }
    }
    rightView(root.Right)
    for i := len(right)-1; i >= 0; i-- {
        out = append(out, right[i])
    }
    return out
}