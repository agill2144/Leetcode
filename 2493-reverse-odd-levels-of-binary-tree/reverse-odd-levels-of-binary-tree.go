/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func reverseOddLevels(root *TreeNode) *TreeNode {
    var dfs func(r1, r2 *TreeNode, level int)
    dfs = func(r1, r2 *TreeNode, level int) {
        // base
        if r1 == nil || r2 == nil {return}

        // logic
        if level % 2 != 0 {
            r1.Val, r2.Val = r2.Val, r1.Val
        }
        dfs(r1.Left, r2.Right, level+1)
        dfs(r1.Right, r2.Left, level+1)
    }
    dfs(root.Left, root.Right, 1)
    return root
}

// bfs: collect odd level nodes while processing that level
// then reverse those node values after that level is over
// func reverseOddLevels(root *TreeNode) *TreeNode {
//     if root == nil {return root}
//     q := []*TreeNode{root}
//     level := 0
//     for len(q) != 0 {
//         qSize := len(q)
//         levelNodes := []*TreeNode{}
//         for qSize != 0 {
//             dq := q[0]
//             q = q[1:]
//             if level % 2 != 0 {levelNodes = append(levelNodes, dq)}
//             if dq.Left != nil {q = append(q, dq.Left)}
//             if dq.Right != nil {q = append(q, dq.Right)}
//             qSize--
//         }
//         for i := 0; i < len(levelNodes)/2; i++ {
//             levelNodes[i].Val, levelNodes[len(levelNodes)-1-i].Val = levelNodes[len(levelNodes)-1-i].Val, levelNodes[i].Val
//         }
//         level++
//     }
//     return root
// }