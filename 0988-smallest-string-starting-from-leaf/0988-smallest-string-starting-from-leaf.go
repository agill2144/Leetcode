/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func smallestFromLeaf(root *TreeNode) string {
    smallest := ""
    var dfs func(path string, r *TreeNode)
    dfs = func(path string, r *TreeNode){
        // base
        if r == nil {return}
        
        // logic
        path += string('a'+r.Val)
        if r.Left == nil && r.Right == nil {
            rev := reverse(path)
            //  -1 if a < b
            if smallest == "" || strings.Compare(rev, smallest) == -1 {
                smallest = rev
            }
            return
        }
        dfs(path, r.Left)
        dfs(path, r.Right)
    }
    dfs("", root)
    return smallest
}

func reverse(a string) string{
    out := ""
    for i := len(a)-1; i >= 0; i-- {out += string(a[i])}
    return out
}