/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func flatten(root *TreeNode) {
    if root == nil {return}
    curr := root
    for curr != nil {
        if curr.Left == nil {
            curr = curr.Right
        } else {
            // find the right-tail of left child
            tail := curr.Left
            for tail.Right != nil {
                tail = tail.Right
            }
            // connect tail to curr right
            tail.Right = curr.Right
            // bring curr's left child to its right side
            curr.Right = curr.Left
            curr.Left = nil
            curr = curr.Right
        }
    }
}

// func flatten(root *TreeNode)  {
//     var dfs func(r *TreeNode) *TreeNode
//     dfs = func(r *TreeNode) *TreeNode {
//         // base
//         if r == nil {return nil}

//         // logic
//         lt := dfs(r.Left)
//         rt := dfs(r.Right)
//         lh := r.Left
//         rh := r.Right
//         if lh != nil {
//             r.Right = lh
//             lt.Right = rh
//             r.Left = nil
//         }
//         if rt != nil {return rt}
//         if lt != nil {return lt}
//         return r
//     }
//     dfs(root)

// }