/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


// approach : traverse with bst constraints 
// time: o(n) 
// space: o(h)
func rangeSumBST(root *TreeNode, low int, high int) int {
    sum := 0
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {
            return
        }
        // preorder
        // if r.Val >= low && r.Val <= high {
        //     sum+=r.Val
        // }
        
        if r.Val > low {
            dfs(r.Left)
        }
        
        // inorder
        if r.Val >= low && r.Val <= high {
            sum+=r.Val
        }
        
        if r.Val < high {
            dfs(r.Right)
        }
        
    }
    dfs(root)
    return sum
}


// plain postOrder traversal
// time: o(n)
// space: o(h)
// func rangeSumBST(root *TreeNode, low int, high int) int {
//     sum := 0
//     var dfs func(r *TreeNode)
//     dfs = func(r *TreeNode) {
//         // base
//         if r == nil {
//             return
//         }
       
//         // logic
//         dfs(r.Left)
//         dfs(r.Right)
//         if r.Val >= low && r.Val <= high {
//             sum+=r.Val
//         }
//     }
//     dfs(root)
//     return sum
// }

// plain preorder traversal
// time: o(n)
// space: o(h)
// func rangeSumBST(root *TreeNode, low int, high int) int {
//     sum := 0
//     var dfs func(r *TreeNode)
//     dfs = func(r *TreeNode) {
//         // base
//         if r == nil {
//             return
//         }
       
//         if r.Val >= low && r.Val <= high {
//             sum+=r.Val
//         }
//         // logic
//         dfs(r.Left)
//         dfs(r.Right)
//     }
//     dfs(root)
//     return sum
// }


// plain inorder traversal
// time: o(n)
// space: o(h)
// func rangeSumBST(root *TreeNode, low int, high int) int {
//     sum := 0
//     var dfs func(r *TreeNode)
//     dfs = func(r *TreeNode) {
//         // base
//         if r == nil {
//             return
//         }
//         // logic
//         dfs(r.Left)
//         if r.Val >= low && r.Val <= high {
//             sum+=r.Val
//         }

//         dfs(r.Right)
//     }
//     dfs(root)
//     return sum
// }