/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// iterative using BFS
// time: o(n)
// space: o(maxWidthOfTheTree)
// func levelOrder(root *TreeNode) [][]int {
//     if root == nil {
//         return nil
//     }
//     var result [][]int
//     q := []*TreeNode{root}
//     for len(q) != 0 {
//         qSize := len(q)
//         level := []int{}
//         for qSize != 0 {
//             dq := q[0]
//             q = q[1:]
//             level = append(level, dq.Val)
//             if dq.Left != nil {q = append(q, dq.Left)}
//             if dq.Right != nil {q = append(q, dq.Right)}
//             qSize--
//         }
//         result = append(result, level)
//     }
//     return result
// }



// approach : level order using DFS ( by maintaining a level at each node in recursion stack )
// level order can be done using DFS as well
// At each node in the recursion stack we will maintain a level -- which represents the idx in the nested array
// if the len of the result array == level - it means result[level] idx does not exist - for it to exist, the len > level.
// therefore we will initialize a new array within, so that result[level] exists.
// once the result[level] array does exist, we will simply append the current level root.Val to result[level]

// time: o(n)
// space: o(h) -- at worse, we will have the max height number of nodes in our recursion stack


func levelOrder(root *TreeNode) [][]int {
    
    var result [][]int
    var dfs func(a *TreeNode, level int)
    dfs = func(a *TreeNode, level int) {
        // base
        if a == nil {return}
        
        // logic
        if len(result) == level {
            result = append(result, []int{})
        }
        result[level] = append(result[level], a.Val)
        dfs(a.Left, level+1)
        dfs(a.Right, level+1)
    }
    
    dfs(root, 0)
    return result
    
}