/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    out := []int{}    
    var dfs func(r *TreeNode, level int)
    dfs = func(r *TreeNode, level int) {
        // base
        if r == nil {return}
        
        // logic
        if level == len(out) {
            out = append(out, r.Val)
        }
        dfs(r.Right, level+1)
        dfs(r.Left, level+1)
    }
    dfs(root, 0)
    return out
}

/*
    approach: level order using BFS
    - save the last item of each level
        - when qSize == 1, this is the last element of a level
        - save it
    time = o(n)
    space = o(maxWidth)
    - at one point, the max width of a level will be added to our queue
*/
// func rightSideView(root *TreeNode) []int {
//     if root == nil {return nil}
//     out := []int{}
//     q := []*TreeNode{root}
//     for len(q) != 0 {
//         qSize := len(q)
//         for qSize != 0 {
//             dq := q[0]
//             q = q[1:]
//             if qSize == 1 {
//                 // this is the last element in this level
//                 out = append(out, dq.Val)
//             }
//             if dq.Left != nil {q = append(q, dq.Left)}
//             if dq.Right != nil {q = append(q, dq.Right)}
//             qSize--
//         }
//     }
//     return out
// }