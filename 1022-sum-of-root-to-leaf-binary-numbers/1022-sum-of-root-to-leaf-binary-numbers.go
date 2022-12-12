/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) int {
    var total int64
    var dfs func(r *TreeNode, path string)
    dfs = func(r *TreeNode, path string) {
        // base
        if r == nil {return}
        
        // logic
        path += fmt.Sprintf("%v", r.Val)
        if r.Left == nil && r.Right == nil {
            n, _ := strconv.ParseInt(path, 2, 64)
            total += n
            return
        }
        dfs(r.Left, path)
        dfs(r.Right, path)
    }
    dfs(root, "")
    return int(total)
}