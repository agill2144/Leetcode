/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // same as: https://leetcode.com/problems/lowest-common-ancestor-of-deepest-leaves/description/
 // lca of ALL nodes at a particular level/depth
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
    maxDepth := math.MinInt64     
    var dfs func(r *TreeNode, depth int)
    dfs = func(r *TreeNode, depth int) {
        // base
        if r == nil {return}

        // logic
        maxDepth = max(maxDepth, depth)
        dfs(r.Left, depth+1)
        dfs(r.Right, depth+1)
    }
    dfs(root, 0)
    var lca func(r *TreeNode, depth int) *TreeNode
    lca = func(r *TreeNode, depth int) *TreeNode {
        // base
        if r == nil {return nil}

        // logic
        if depth == maxDepth {
            return r
        }
        left := lca(r.Left, depth+1)
        right := lca(r.Right, depth+1)
        if left != nil && right != nil {
            return r
        }
        if left != nil {return left}
        return right
    }
    return lca(root,0)
}