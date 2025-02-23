/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
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
        if ptr == len(pre) {return nil}
        if currDepth != 0 && currDepth > pre[ptr].depth {return nil}

        // logic
        root := &TreeNode{Val:pre[ptr].val}
        ptr++
        root.Left = dfs(currDepth+1)
        root.Right = dfs(currDepth+1)
        return root
    }
    return dfs(0)
    
}