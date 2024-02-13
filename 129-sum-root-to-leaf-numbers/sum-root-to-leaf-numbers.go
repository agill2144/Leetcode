/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// we can maintain the same path in global var
// and just backtrack once a node is fully processed ( i.e left and right child are done )
func sumNumbers(root *TreeNode) int {
    total := 0
    path := 0
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}

        // logic
        // action
        path = (path*10)+r.Val
        
        // recurse
        dfs(r.Left)
        if r.Left == nil && r.Right == nil {
            total += path
        }
        dfs(r.Right)
        
        // backtrack - i.e remove the last digit
        path = path/10
    }
    dfs(root)
    return total
}

// maintain path as part of recursion
// once we are at the leaf node, add the path to total sum
// recursion will take the path back to its original value when
// the recursion goes back to a parent node
// no need to backtrack anything because we are not using a reference data type
// func sumNumbers(root *TreeNode) int {
//     total := 0
//     var dfs func(r *TreeNode, p int)
//     dfs = func(r *TreeNode, p int) {
//         // base
//         if r == nil {return}
        
//         // logic
//         p = p * 10 + r.Val
//         dfs(r.Left, p)
//         dfs(r.Right, p)
//         if r.Left == nil && r.Right == nil {
//             total += p
//         }   
//     }
//     dfs(root, 0)
//     return total
// }