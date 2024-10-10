/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
    maxDepth := 0
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
    lca = func(r *TreeNode,  depth int) *TreeNode {
        // base
        if r == nil {return r}
        // logic
        left := lca(r.Left, depth+1)
        right := lca(r.Right, depth+1)
        if depth == maxDepth {return r}
        if left != nil && right != nil {return r}
        if left != nil {return left}
        return right
         
    }
    return lca(root, 0)
}