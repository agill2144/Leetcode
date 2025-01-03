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
    out := []int{root.Val}
    left := func(r *TreeNode) {
        for r != nil {
            if r.Left == nil && r.Right == nil {return}
            out = append(out, r.Val)
            if r.Left != nil {
                r = r.Left
            } else {
                r = r.Right
            }
        }
    }
    left(root.Left)
    // leaves from left to right
    var leaves func(r *TreeNode)
    leaves = func(r *TreeNode) {
        // base
        if r == nil {return}

        // logic
        if r != root && r.Left == nil && r.Right == nil {
            out = append(out, r.Val)
            return
        }
        leaves(r.Left)
        leaves(r.Right)
    }
    leaves(root)

    // right side
    // but we need it in reverse order
    // therefore keeping track of leftPtr
    leftPtr := len(out)
    right := func(r *TreeNode) {
        for r != nil {
            if r.Left == nil && r.Right == nil {return}
            out = append(out, r.Val)
            if r.Right != nil {
                r = r.Right
            } else {
                r = r.Left
            }
        }
    }
    right(root.Right)
    rightPtr := len(out)-1
    for leftPtr < rightPtr {
        out[leftPtr], out[rightPtr] = out[rightPtr], out[leftPtr]
        leftPtr++
        rightPtr--
    }
    return out
}