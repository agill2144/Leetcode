/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
    set := map[int]struct{}{}
    for _, ele := range to_delete {set[ele] = struct{}{}}

    out := []*TreeNode{}
    var dfs func(r *TreeNode, parent *TreeNode)
    dfs = func(r *TreeNode, parent *TreeNode) {
        // base
        if r == nil {return}

        // logic
        dfs(r.Left, r)
        dfs(r.Right, r)
    
        if _, ok := set[r.Val]; ok {
            if r.Left != nil {out = append(out, r.Left)}
            if r.Right != nil {out = append(out, r.Right)}
            if parent != nil {
                if parent.Left == r {parent.Left = nil}
                if parent.Right == r {parent.Right = nil}
            }
        }
    }

    dfs(root, nil)
    if _, ok := set[root.Val]; !ok {out = append(out, root)}
    return out

}