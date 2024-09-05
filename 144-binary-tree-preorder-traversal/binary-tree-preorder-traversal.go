/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// iterative
func preorderTraversal(root *TreeNode) []int {
    st := []*TreeNode{}
    out := []int{}
    for root != nil || len(st) != 0 {
        for root != nil {
            out = append(out, root.Val)
            st = append(st, root)
            root = root.Left
        }
        root = st[len(st)-1]
        st = st[:len(st)-1]
        root = root.Right
    }
    return out
}

// dfs
// func preorderTraversal(root *TreeNode) []int {
//     out := []int{}
//     var dfs func(r *TreeNode)
//     dfs = func(r *TreeNode) {
//         // base
//         if r == nil {return}
//         // logic
//         out = append(out, r.Val)
//         dfs(r.Left)
//         dfs(r.Right)
//     }
//     dfs(root)
//     return out
// }