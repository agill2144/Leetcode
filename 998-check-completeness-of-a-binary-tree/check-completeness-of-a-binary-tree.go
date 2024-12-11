/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    approach: dfs; map nodes to idxs in an array
    - we would need to know how many nodes exists
        - we can run a vanilla dfs to get this count
    - heaps also have a tree like structure, but they are stored in an array
    - in heaps, we would map a parents left child to idx : 2*i+1
    - in heaps, we would map a parents right child to idx : 2*i+2
    - we can map each node in the tree to an idx
    - root node is at idx 0, therefore we start i = 0
    - then left child is at 2*i+1
    - and right child is at 2*i+2
    - just like ^ heap childs in an array
    - how does this help?
    - we know how many idxs are supposed to exists ( count-1 )
        - we have this count from running a dfs earlier
    - if an idx of a node comes out to be >= count, this cannot be mapped to an array of size count
    - and this can only happen
        - if the tree was incomplete
        - as in, the node went out of bounds, could have been written to the left side side
        - but this node was on the right side, skipping a empty left spot
        - this node was not as far left as possible
        - or
        - this node was on the last level but the prev level had space!
    tc = o(2n)
    sc = o(2n)
*/
func isCompleteTree(root *TreeNode) bool {
    if root == nil {return true}
    var count func(r *TreeNode) int
    count = func(r *TreeNode) int {
        // base
        if r == nil {return 0}
        // logic
        left := count(r.Left)
        right := count(r.Right)
        return left+right+1
    }
    n := count(root)
    var dfs func(r *TreeNode, idx int) bool
    dfs = func(r *TreeNode, idx int) bool {
        // base
        if r == nil {return true}

        // logic
        if idx >= n {return false}
        if !dfs(r.Left, 2*idx+1) {return false}
        return dfs(r.Right, 2*idx+2)
    }
    return dfs(root, 0)
}

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
// func isCompleteTree(root *TreeNode) bool {
//     if root == nil {return true}
//     q := []*TreeNode{root}    
//     seenNil := false
//     for len(q) != 0 {
//         dq := q[0]
//         q = q[1:]
//         if dq == nil {seenNil = true}
//         if dq != nil {
//             if seenNil {return false}
//             q = append(q, dq.Left)
//             q = append(q, dq.Right)
//         }
//     }
//     return true
// }