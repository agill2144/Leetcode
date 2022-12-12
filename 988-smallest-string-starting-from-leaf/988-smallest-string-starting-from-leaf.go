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
    var dfs func(r *TreeNode, path string)
    dfs = func(r *TreeNode, path string) {
        // base
        if r == nil {return}
        
        // logic
        path += string('a'+r.Val)
        if r.Left == nil && r.Right == nil {
            reversed := rev(path)
            val := strings.Compare(reversed, smallest)
            if smallest == "" || val == -1 {
                smallest = reversed
            }
            return
        }
        dfs(r.Left, path)
        dfs(r.Right, path)
    }
    dfs(root, "")
    return smallest
}

func rev(s string) string {
    out := new(strings.Builder)
    for i := len(s)-1; i >= 0; i-- {
        out.WriteByte(s[i])
    }
    return out.String()
}