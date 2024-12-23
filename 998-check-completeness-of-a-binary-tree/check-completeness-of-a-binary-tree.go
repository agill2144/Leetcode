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
    - another way to look at it is by doing level order
    - once we run into a null node while processing a level
    - there shouldnt be any more real nodes after the null node for that level
        - and for the next level tooooo
        - next level should not have any more non-null nodes
*/
func isCompleteTree(root *TreeNode) bool {
    if root == nil {return true}
    q := []*TreeNode{root}
    seenNil := false
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        if dq == nil {
            seenNil = true
        } else {
            if seenNil {return false}
            q = append(q, dq.Left)
            q = append(q, dq.Right)
        }
    }
    return true
}


/*
    approach: dfs with heap idxing
    - complete tree , consider heap idxs
    - heap idxs will mark nodes with an idx
    - we are essentially mapping each node to an array using heap idxs
    - heaps left child idx in the array would be; 2*i+1
    - heaps right child idx in the array would be; 2*i+2
    - similarly we will start with 0 as the initial idx
    - then for left child, the idx will be 2*i+1 and for right child ; 2*i+2
    - how does this help? we dont have an array to physically map node to their idx positions
    - we dont need a physical array
    - if we knew the size of this array, then we would know the last possible idx
    - array will store all n nodes of tree
    - therefore we can perform a dfs to count nodes of the tree
    - now we know the size of the array, we dont need to create it
    - if a heap idx goes out of bounds ( idx >= n ), it means we do not have a complete tree
    - heap idxs increase by 1 from top to bottom left to right
    - and if we jumped over a nil node and went to right node , that right node idx will be greater than last possible idx
    - therefore not a complete tree
    tc = o(2n)
    sc = o(2n)
*/
// func isCompleteTree(root *TreeNode) bool {
//     var count func(r *TreeNode) int
//     count = func(r *TreeNode) int {
//         // base
//         if r == nil {return 0}
//         // logic
//         return 1 + count(r.Left) + count(r.Right)
//     }
//     n := count(root)
//     if n == 0 {return true}
//     var dfs func(r *TreeNode, idx int) bool
//     dfs = func(r *TreeNode, idx int)bool {
//         // base
//         if r == nil {return true}

//         // logic
//         if idx >= n {return false}
//         if !dfs(r.Left, 2*idx+1) {return false}
//         return dfs(r.Right, 2*idx+2)
//     }
//     return dfs(root, 0)
// }