/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
    if root == nil {return nil}
    toDelete := map[int]bool{}
    for i := 0; i < len(to_delete); i++ {toDelete[to_delete[i]] = true}
    roots := []*TreeNode{}
    var dfs func(r, p *TreeNode)
    dfs = func(r, p *TreeNode) {
        // base
        if r == nil {return}

        // logic
        dfs(r.Left, r)
        dfs(r.Right, r)
        if toDelete[r.Val] {
            if r.Left != nil {roots = append(roots, r.Left)}
            if r.Right != nil {roots = append(roots, r.Right)}
            if p != nil && p.Left == r {p.Left = nil}
            if p != nil && p.Right == r {p.Right = nil} 
        }
    }
    dfs(root, nil)
    if !toDelete[root.Val] {
        roots = append(roots, root)
    }
    return roots
}