/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
    if depth == 1 {
        newRoot := &TreeNode{Val: val}
        newRoot.Left = root
        return newRoot
    }
    var dfs func(r *TreeNode, currDepth int)
    dfs = func(r *TreeNode, currDepth int) {
        // base
        if r == nil {return}

        // logic
        if currDepth+1 == depth {
            currLeft := r.Left
            currRight := r.Right
            r.Left = &TreeNode{Val: val}
            r.Right = &TreeNode{Val: val}
            r.Left.Left = currLeft
            r.Right.Right = currRight
        }
        dfs(r.Left, currDepth+1)
        dfs(r.Right, currDepth+1)
    }
    dfs(root, 1)
    return root
}