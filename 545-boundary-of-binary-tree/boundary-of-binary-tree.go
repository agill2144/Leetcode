/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func boundaryOfBinaryTree(root *TreeNode) []int {
    out := []int{root.Val}
    left := func(r *TreeNode) {
        for r != nil {
            if r.Left == nil && r.Right == nil {break}
            out = append(out, r.Val)
            if r.Left != nil {r = r.Left} else {r = r.Right} 
        }
    }
    left(root.Left)
    // leaves
    var leaves func(r *TreeNode) 
    leaves = func(r *TreeNode) {
        // base
        if r == nil {return}

        // logic
        leaves(r.Left)
        leaves(r.Right)
        if r.Left == nil && r.Right == nil && r != root {
            out = append(out, r.Val)
        }
    }
    leaves(root)

    // right
    startIdx := len(out)
    right := func(r *TreeNode) {
        for r != nil {
            if r.Left == nil && r.Right == nil {break}
            out = append(out, r.Val)
            if r.Right != nil {r = r.Right} else {r = r.Left}            
        }
    }
    right(root.Right)
    l, r := startIdx, len(out)-1
    for l < r {
        out[l], out[r] = out[r], out[l]
        l++; r--
    }
    return out
}