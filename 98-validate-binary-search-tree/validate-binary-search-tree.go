/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // inorder on a BST = sorted order
func isValidBST(root *TreeNode) bool {
    var dfs func(r *TreeNode, min, max int) bool
    dfs = func(r *TreeNode, min, max int) bool {
        // base
        if r == nil {return true}

        // logic
        if r.Val <= min || r.Val >= max {return false}
        if !dfs(r.Left, min, r.Val) {return false}
        if !dfs(r.Right, r.Val, max) {return false}
        return true
    }
    return dfs(root, math.MinInt64, math.MaxInt64)
}
// func isValidBST(root *TreeNode) bool {
//     var prev *TreeNode
//     var dfs func(r *TreeNode) bool
//     dfs = func(r *TreeNode) bool {
//         // base
//         if r == nil {return true}

//         // logic
//         left := dfs(r.Left)
//         if !left {return false}
//         if prev != nil {
//             if prev.Val >= r.Val {return false}
//         }
//         prev = r
//         return dfs(r.Right)
//     }
//     return dfs(root)
// }