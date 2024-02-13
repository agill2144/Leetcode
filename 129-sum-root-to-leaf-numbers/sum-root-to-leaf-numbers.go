/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    approach: dfs without global sum ( bottom up - childs returning to parent )
    - maintain path as part of recursion
    - once we are at the leaf node, return that path to its parent
    - recursion will take the path back to its original value when
    - the recursion goes back to a parent node    
    time: o(n)
    space: o(n)
*/
func sumNumbers(root *TreeNode) int {
    var dfs func(r *TreeNode, path int) int
    dfs = func(r *TreeNode, path int) int {
        // base
        if r == nil {return 0}

        // logic
        path = (path * 10) + r.Val
        left := dfs(r.Left, path)
        right := dfs(r.Right, path)
        if r.Left == nil && r.Right == nil {return path}
        return left+right
    }
    return dfs(root, 0)
}

/*
    approach: dfs with global sum
    - maintain path as part of recursion
    - once we are at the leaf node, add the path to total sum
    - recursion will take the path back to its original value when
    - the recursion goes back to a parent node
    - no need to backtrack anything because we are not using a reference data type
    
    time: o(n)
    space: o(n)

*/
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



// we can maintain the same path in global var
// and just backtrack once a node is fully processed ( i.e left and right child are done )
// func sumNumbers(root *TreeNode) int {
//     total := 0
//     path := 0
//     var dfs func(r *TreeNode)
//     dfs = func(r *TreeNode) {
//         // base
//         if r == nil {return}

//         // logic
//         // action
//         path = (path*10)+r.Val
//         // recurse
//         dfs(r.Left)
//         dfs(r.Right)
//         if r.Left == nil && r.Right == nil {
//             total += path
//         }
//         // backtrack - i.e remove the last digit
//         path = path/10
//     }
//     dfs(root)
//     return total
// }
