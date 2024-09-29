/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    var prev *TreeNode
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}

        // logic
        dfs(r.Right)
        dfs(r.Left)
        r.Right = prev
        r.Left = nil
        prev = r
    }
    dfs(root)
    
}
// func flatten(root *TreeNode)  {
//     if root == nil {return}
//     curr := root
//     for curr != nil {
//         if curr.Left == nil {
//             curr = curr.Right
//         } else {
//             tmp := curr.Left
//             for tmp.Right != nil {
//                 tmp = tmp.Right
//             }
//             tmp.Right = curr.Right
//             curr.Right = curr.Left
//             curr.Left = nil
//         }
//     }
// }