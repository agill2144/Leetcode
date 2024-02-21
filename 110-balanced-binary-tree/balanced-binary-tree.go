/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// height-balanced binary tree is when height diff between left and right child <= 1
// for each node, we can get the heigh of left subtree and right subtree and calc diff
// this will be o(n^2) ; top down solution
// instead we can do this in linear fashing using bottom up
func isBalanced(root *TreeNode) bool {
    var dfs func(r *TreeNode) (int, bool)
    dfs = func(r *TreeNode) (int, bool) {
        // base
        if r == nil {return 0, true}
        // logic
        left, lOk := dfs(r.Left)
        if !lOk {return 0,false}
        right, rOk := dfs(r.Right)
        if !rOk {return 0,false}

        if abs(right-left) > 1 {return 0,false}

        return max(left, right)+1, true
    }
    _, ok := dfs(root)
    return ok
}


func abs(x int) int {
    if x < 0 {return x * -1}
    return x
}