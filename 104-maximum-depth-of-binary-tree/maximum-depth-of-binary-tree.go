/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
    top-down recursion
    - we can keep track of number of nodes in recursion at each node
    - always save the number of nodes when number of nodes in recursion > max seen so far
    time = o(n)
    space = o(n) - skewed tree
*/
func maxDepth(root *TreeNode) int {
    if root == nil {return 0}
    maxN := 0    
    var dfs func(r *TreeNode, count int) 
    dfs = func(r *TreeNode, count int) {
        // base
        if r == nil {
            return 
        }
        // logic
        maxN = max(maxN, count)
        dfs(r.Left, count+1)
        dfs(r.Right, count+1)
    }
    dfs(root, 1)
    return maxN
}