/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    approach: bottom up
    - alternatively, we can count from bottom of the tree
    - that is, once we reach a leaf node, we will start the counter
    - and return back 1 to its parent
    - then the parent will compare which is greater in depth ( left or right return value )
    - then the parent will take the max depth ( since we want the max depth ) and add 1 to it
        - why add 1 to it?
        - because parent node is a also a child node of some other parent
        - meaning we also need to count this node

    time = o(n)
    space = o(n) ; worst case we have a skewed tree
*/
func maxDepth(root *TreeNode) int {
    var dfs func(r *TreeNode) int
    dfs = func(r *TreeNode) int {
        // base
        if r == nil {return 0}

        // logic
        left := dfs(r.Left)
        right := dfs(r.Right)
        return max(left,right)+1
    }
    return dfs(root)
}


 /*
    approach: top down
    - using a global var to keep track of the max depth
    - keep track of the depth ( number of nodes ) as we are traversing the tree
    - once we hit the leaf node, 
    - compare and save the depth with the global var

    time = o(n)
    space = o(maxHeightOfTree) ; worst case its o(n) 
 */
// func maxDepth(root *TreeNode) int {
//     out := 0
//     var dfs func(r *TreeNode, h int)
//     dfs = func(r *TreeNode, h int) {
//         // base
//         if r == nil {return}

//         // logic
//         h++
//         if r.Left == nil && r.Right == nil {
//             out = max(h, out)
//         }
//         dfs(r.Left, h)
//         dfs(r.Right, h)
//     }
//     dfs(root, 0)
//     return out
// }