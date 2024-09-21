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
    curr := root.Left
    for curr != nil {
        if curr.Left == nil && curr.Right == nil {break}
        out = append(out,curr.Val)
        if curr.Left != nil {
            curr = curr.Left
        } else {
            curr = curr.Right
        }
    }

    // collect leaves from left to right
    var leaves func(r *TreeNode)
    leaves = func(r *TreeNode) {
        // base
        if r == nil {return}
        // logic
        if r.Left == nil && r.Right == nil && r != root {out = append(out, r.Val); return}
        leaves(r.Left)
        leaves(r.Right)
    }
    leaves(root)

    // collect right side
    rightStartPtr := len(out)
    curr = root.Right
    for curr != nil {
        // we have reached the leaf node
        if curr.Left == nil && curr.Right == nil {break}
        out = append(out,curr.Val)
        if curr.Right != nil {curr = curr.Right} else {curr = curr.Left}
    }
    end := len(out)-1
    for rightStartPtr < end {
        out[rightStartPtr], out[end] =  out[end],out[rightStartPtr]
        rightStartPtr++
        end--
    }
    return out
}