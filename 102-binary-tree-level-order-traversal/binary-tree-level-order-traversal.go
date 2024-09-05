/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// level order using dfs
// maintain level as dfs arg
// and use the level as a idx in output array
func levelOrder(root *TreeNode) [][]int {
    out := [][]int{}
    var dfs func(r *TreeNode, level int)
    dfs = func(r *TreeNode, level int) {
        // base
        if r == nil {return}

        // logic
        if len(out) == level {
            out = append(out, []int{})
        }
        out[level] = append(out[level], r.Val)
        dfs(r.Left, level+1)
        dfs(r.Right, level+1)
    }
    dfs(root, 0)
    return out
}


// level order using bfs
// func levelOrder(root *TreeNode) [][]int {
//     out := [][]int{}
//     if root == nil {return out}
//     q := []*TreeNode{root}    
//     for len(q) != 0 {
//         qSize := len(q)
//         level := []int{}
//         for qSize != 0 {
//             dq := q[0]
//             q = q[1:]
//             level = append(level, dq.Val)
//             if dq.Left != nil {
//                 q = append(q, dq.Left)
//             }
//             if dq.Right != nil {
//                 q = append(q, dq.Right)
//             }
//             qSize--
//         }
//         out = append(out, level)
//     }
//     return out
// }