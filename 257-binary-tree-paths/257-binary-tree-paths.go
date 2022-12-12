/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
    result := []string{}
    var dfs func(path string, r *TreeNode)
    dfs = func(path string, r *TreeNode) {
        // base
        if r == nil {
            return
        }
        
        // logic
        path += fmt.Sprintf("%v",r.Val)
        if r.Left == nil && r.Right == nil {
            result = append(result, path)
            return
        }
        if r.Left != nil || r.Right != nil {path+="->"}
        dfs(path, r.Left)
        dfs(path, r.Right)
    }
    dfs("", root)
    return result
}