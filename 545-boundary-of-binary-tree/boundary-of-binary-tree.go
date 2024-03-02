/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    approach: brute force
    - collect left view ( without leaves )
    - collect ONLY leaves
    - collect right view ( without leaves )
        - but then reverse the right view
        - therefore collect right view in a stack
        - and process stack from top ( i.e reverse order of right view ; from bottom to top )
    
    time = o(n)
    space = o(n)
*/
func boundaryOfBinaryTree(root *TreeNode) []int {
    out := []int{}
    if !isLeaf(root) {
        out = append(out, root.Val)
    }

    // collect left view
    l := root.Left
    for l != nil {
        if !isLeaf(l) {
            out = append(out, l.Val)
        }
        if l.Left != nil {
            l = l.Left
        } else {
            l = l.Right
        }
    }
    // collect leaves
    var leaves func(r *TreeNode)
    leaves = func(r *TreeNode) {
        // base
        if r == nil {return}

        // logic
        if isLeaf(r) {
            out = append(out, r.Val)
        }        
        leaves(r.Left)
        leaves(r.Right)

    }
    leaves(root)

    rightSt := []int{}
    // collect right view
    r := root.Right
    for r != nil {
        if !isLeaf(r) {
            rightSt = append(rightSt, r.Val)
        }
        if r.Right != nil {
            r = r.Right
        } else {
            r = r.Left
        }
    }
    for len(rightSt) != 0 {
        out = append(out, rightSt[len(rightSt)-1])
        rightSt = rightSt[:len(rightSt)-1]
    }
    
    return out
    
}

func isLeaf(r *TreeNode) bool {
    return r.Left == nil && r.Right == nil
}