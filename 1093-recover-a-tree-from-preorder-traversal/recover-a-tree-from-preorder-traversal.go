/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// num of dashes == depth of node
// cannot add left/right child if current depth > current node's depth
// traversal is given in a pre-order travsersal form
// parse it , and create a list out of [nodeval, depth]
// then use dfs to construct it from the list
// root node is always at idx 0
// therefore keep a ptr to track what idx we are on
// also keep track of current depth in dfs (starting from 0)
// if depth at currDepth in dfs > list[ptr] depth, this node cannot be a child
// of a prev parent node we just created, therefore return nil
func recoverFromPreorder(traversal string) *TreeNode {
    type node struct {
        depth int
        val int
    }
    pre := []node{}
    depth := 0
    for i := 0; i < len(traversal); i++ {
        if traversal[i] == '-' {
            depth++
        } else {
            val := 0
            for i < len(traversal) && traversal[i] >= '0' && traversal[i] <= '9' {
                val = val * 10 + int(traversal[i]-'0')
                i++
            }
            pre = append(pre, node{depth, val})
            depth = 1
        }
    }
    ptr := 0
    var dfs func(currDepth int) *TreeNode
    dfs = func(currDepth int) *TreeNode {
        // base
        if ptr == len(pre) || currDepth > pre[ptr].depth {return nil}

        // logic
        root := &TreeNode{Val:pre[ptr].val}
        ptr++
        root.Left = dfs(currDepth+1)
        root.Right = dfs(currDepth+1)
        return root
    }
    return dfs(0)
    
}