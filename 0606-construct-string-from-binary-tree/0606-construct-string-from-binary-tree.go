/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(root *TreeNode) string {
    out := new(strings.Builder)
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}        
        // logic
        out.WriteString(fmt.Sprintf("%v", r.Val))
        if r.Left == nil && r.Right == nil {return}
        

        out.WriteString("(")
        dfs(r.Left)
        out.WriteString(")")
        
        if r.Right != nil {
            out.WriteString("(")
            dfs(r.Right)
            out.WriteString(")")
        }
        
        
        
    }
    dfs(root)
    return out.String()
}