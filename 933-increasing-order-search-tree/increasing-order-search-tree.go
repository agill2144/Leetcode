/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
    ans := &TreeNode{Val: 0}
    var dfs func(r, p *TreeNode)
    dfs = func(r, p *TreeNode) {
        // base
        if r == nil {return}

        // logic
        dfs(r.Left, r)
        if ans.Right == nil {ans.Right = r}

        // non-leaf node
        if r.Left != nil  {
            if r.Left != nil {
                t := r.Left
                if p != nil {
                    if p.Right == r {
                        p.Right = t
                    } else if p.Left == r {
                        p.Left= t
                    }
                }
                for t.Right != nil {t = t.Right}
                t.Right = r
                r.Left = nil
            }
        }        

        dfs(r.Right,r)
    }
    dfs(root, nil)
    return ans.Right
}
