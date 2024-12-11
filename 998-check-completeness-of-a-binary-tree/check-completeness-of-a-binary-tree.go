/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    approach: bfs
    - once we run into a null node
    - we should no longer run into a null node
    - for each non-null node, blindly add left and right childs
    - if we run into a non-null node, a flag should tell us whether we had seen null node before this
    - if the flag tells us that we have seen a null node before, and we run into a non null node, this tree is not complete

    tc = o(n)
    sc = o(maxWidth of the tree); o(n/2) ; o(n)
*/
func isCompleteTree(root *TreeNode) bool {
    if root == nil {return true}
    q := []*TreeNode{root}    
    seenNil := false
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        if dq == nil {seenNil = true}
        if dq != nil {
            if seenNil {return false}
            q = append(q, dq.Left)
            q = append(q, dq.Right)
        }
    }
    return true
}