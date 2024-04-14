/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// a node X in the tree is named good if in the path from root to X there are no nodes with a value greater than X
// meaning if we have seen a node > X val, its not a good node
// so lets just maintain a max seen so far as we go down a path
func goodNodes(root *TreeNode) int {
    count := 0
    var dfs func(r *TreeNode, maxVal int)
    dfs = func(r *TreeNode, maxVal int) {
        // base
        if r == nil {return}

        // logic
        // found a good node
        if maxVal <= r.Val {count++}
        dfs(r.Left, max(maxVal, r.Val))
        dfs(r.Right, max(maxVal, r.Val))
    }
    dfs(root, math.MinInt64)
    return count
}