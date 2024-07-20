/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
    - find depeest depth first
    - there could be 2 leaves of n/2 leaves (perfect,complete tree)

    - now we have maxDepth, lets assume worst case that all leaves share the same max depth
    - meaning we have n/2 leave at same $maxDepth
    - now we have to find lca of ALL n/2 leaves; "return the lowest common ancestor of its deepest leaves"
    - we know how to get lca of 2 nodes
    - instead of finding lca of all n/2 nodes, can we instead look for lca of a particular depth ? 
    - if we have reached $maxDepth, assume these nodes at $maxDepth are our p and q in lca
    - so now we simply imply that if we have reached $maxDepth, we need to return this node
    - now the traditional lca
        - if the parent gets nodes from left and right, the parent is the lca
        - if parent gets a node only from left side, left is lca
        - if parent gets a node only from right side, right is lca
    - essentially we can find lca of all nodes that are located at a particular level/depth!    

    lca pattern; lca of node p and q
    extend lca to now cover: lca of all nodes at a level/depth
*/

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
    maxDepth := 0

    // top down dfs and keep track of depth at a node
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
        left := lca(r.Left, depth+1)
        right := lca(r.Right, depth+1)
        if depth == maxDepth {
            // we are at the p or q node of lca, return current node
            return r
        }
        if left != nil && right != nil {
            return r
        }
        if left != nil {return left}
        return right
    }
    return lca(root, 0)
}