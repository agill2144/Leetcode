/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    We just need a helper that takes 2 trees
    and that helper will check whether
    the 2 trees are mirror image or not.
*/

func isSymmetric(root *TreeNode) bool {
    return dfs(root, root)
}

// preorder dfs ( processing first, then left and right )
func dfs(a *TreeNode, b *TreeNode) bool {
    if a == nil && b == nil {
        return true
    }
    if a == nil || b == nil || a.Val != b.Val {return false}
    
    return dfs(a.Left, b.Right) && dfs(a.Right, b.Left)
}

// inorder dfs ( go left, then process, and then go right )
// this is slightly messy, I prefer preorder in this case 
// func dfs(a *TreeNode, b *TreeNode) bool {
//     // base
//     if a == nil && b == nil {
//         return true
//     }
//     if a == nil || b == nil {
//         return false 
//     }
    
//     // logic
//     if left := dfs(a.Left, b.Right); !left {
//         return false // no need to go to next if this is already failing
//     }
//     if a == nil || b == nil || a.Val != b.Val {
//         return false
//     }
//     return dfs(a.Right, b.Left)
    
// }
