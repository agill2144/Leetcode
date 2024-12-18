/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCompleteTree(root *TreeNode) bool {
    var count func(r *TreeNode) int
    count = func(r *TreeNode) int {
        // base
        if r == nil {return 0}

        // logic
        return 1 + count(r.Left) + count(r.Right)
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

// func isCompleteTree(root *TreeNode) bool {
//     if root == nil {return true }
//     seenNil := false
//     q := []*TreeNode{root}
//     for len(q) != 0 {
//         dq := q[0]
//         q = q[1:]
//         if dq == nil {
//             seenNil = true
//         } else {
//             if seenNil {return false}
//             q = append(q, dq.Left)
//             q = append(q, dq.Right)
//         }
//     }
//     return true
// }