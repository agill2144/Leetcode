/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
    approach: level using using dfs ( using level as idx )
    - maintain a level as part of recursion ( level or depth or height ) - in a top down manner
        - i.e each child we go to, we increment the level arg by 1
    - if our output list size == level, it means it DOES NOT have a list at $level-idx
        - therefore if output list size == level, append a new empty list
    - otherwise just append to nested level-idx list

    time = o(n)
    space = o(maxHeightOfTree)
*/
func levelOrder(root *TreeNode) [][]int {
    out := [][]int{}
    var dfs func(r *TreeNode, level int)
    dfs = func(r *TreeNode, level int) {
        // base
        if r == nil {return}

        // logic
        if len(out) == level {
            // we do not have a list at idx $level
            out = append(out, []int{})
            // now we do
        }
        out[level] = append(out[level], r.Val)
        dfs(r.Left, level+1)
        dfs(r.Right, level+1)
    }
    dfs(root, 0)
    return out
}

/*
    approah: level order using bfs
    - level order using

    time = o(n)
    space = o(maxWidth of tree) ; i.e last level ; i.e o(n/2); i.e o(n)
*/
// func levelOrder(root *TreeNode) [][]int {
//     if root == nil {return nil}
//     out := [][]int{}
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
//         out = append(out,level)
//     }
//     return out
// }