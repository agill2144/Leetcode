/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
    out := [][]int{}
    var dfs func(r *TreeNode, path []int, sum int)
    dfs = func(r *TreeNode, path []int, sum int) {
        // base
        if r == nil {return}
        
        // logic
        sum += r.Val
        path = append(path, r.Val)
        dfs(r.Left, path, sum)
        
        if r.Left == nil && r.Right == nil && sum == targetSum{
            newL := make([]int, len(path))
            copy(newL, path)
            out = append(out, newL)
            path = path[:len(path)-1]
            return
        }

        
        dfs(r.Right, path, sum)
        path = path[:len(path)-1]
    }
    dfs(root, nil,0)
    return out
}