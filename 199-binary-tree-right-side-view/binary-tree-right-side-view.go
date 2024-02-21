/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int{
    out := []int{}
    var dfs func(r *TreeNode, depth int)
    dfs = func(r *TreeNode, depth int) {
        // base
        if r== nil {return}
        // logic
        if len(out) == depth {
            out = append(out, r.Val)
        }
        dfs(r.Right, depth+1)
        dfs(r.Left, depth+1)
    }
    dfs(root, 0)
    return out
}
// time = o(n)
// space = o(maxWidth)
// func rightSideView(root *TreeNode) []int {
//     if root == nil {return nil}
//     out := []int{}
//     q := []*TreeNode{root}
//     for len(q) != 0 {
//         qSize := len(q)
//         for qSize != 0 {
//             dq := q[0]
//             q = q[1:]
//             qSize--
//             if qSize == 0 {
//                 // this was the last element in the level
//                 out = append(out, dq.Val)
//             }
//             if dq.Left != nil {q = append(q, dq.Left)}
//             if dq.Right != nil {q = append(q, dq.Right)}
//         }
//     }
//     return out
// }