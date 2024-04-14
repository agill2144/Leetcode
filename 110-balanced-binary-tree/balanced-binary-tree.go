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
// if top-down recursion leads to n^2, try bottom up recursion
// i.e go down first, process the bottom of the tree first, return something back to parent
// instead we can do this in linear time using bottom up
func isBalanced(root *TreeNode) bool {
    // height diff <= 1 ( maxLeftHeight-maxRightHeight <= 1 == balanced)
    var dfs func(r *TreeNode) (int, bool)
    dfs = func(r *TreeNode) (int, bool){
        // base
        if r == nil {return 0, true}
        // logic
        left, lok := dfs(r.Left)
        if !lok {return -1, false}
        right, rok := dfs(r.Right)
        if !rok {return -1, false}

        // not balanced when height diff is > 1
        if abs(left-right) > 1 {return -1, false}

        // return the max between left and right to its parent
        return max(left, right)+1, true
    }
    _, ok := dfs(root)
    return ok
}

func abs(x int) int {
    if x < 0 {return x*-1}
    return x
}