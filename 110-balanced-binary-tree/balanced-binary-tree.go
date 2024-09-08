/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*

A height-balanced binary tree is a binary tree in which 
the depth of the two subtrees of every node never 
differs by more than one.

approach: bottom up dfs ( height )
- once left and right child has returned their heights to parent
- parent will decide if so far, childs under me are balanced or not
- if yes, take the max height + 1 ( height logic )
*/
func isBalanced(root *TreeNode) bool {

    var dfs func(r *TreeNode) (int, bool)
    dfs = func(r *TreeNode) (int, bool) {
        // base
        if r == nil {return 0, true}

        // logic
        left, lok := dfs(r.Left)
        right, rok := dfs(r.Right)
        if !lok || !rok {return -1, false}
        if abs(right-left) > 1 {return -1, false}
        return max(left, right)+1, true
    } 
    _,  ok := dfs(root)
    return ok
}

func abs(x int) int {
    if x < 0 {return x*-1}
    return x
}