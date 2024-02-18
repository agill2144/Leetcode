/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
    nodeDepths := map[int]int{}
    maxDepth := math.MinInt64     
    var dfs func(r *TreeNode, depth int)
    dfs = func(r *TreeNode, depth int) {
        // base
        if r == nil {return}

        // logic
        nodeDepths[r.Val] = depth
        maxDepth = max(maxDepth, depth)
        dfs(r.Left, depth+1)
        dfs(r.Right, depth+1)
    }
    dfs(root, 0)
    var lca func(r *TreeNode) *TreeNode
    lca = func(r *TreeNode) *TreeNode {
        // base
        if r == nil {return nil}

        // logic
        if nodeDepths[r.Val] == maxDepth {
            return r
        }
        left := lca(r.Left)
        right := lca(r.Right)
        if left != nil && right != nil {
            return r
        }
        if left != nil {return left}
        return right
    }
    return lca(root)
}