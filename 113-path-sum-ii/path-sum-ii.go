/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, target int) [][]int {
    out := [][]int{}
    var dfs func(r *TreeNode, sum int, path []int)
    dfs = func(r *TreeNode, sum int, path []int) {
        // base
        if r == nil {return}

        // logic
        sum += r.Val
        path = append(path, r.Val)
        if r.Left == nil && r.Right == nil {
            if sum == target {
                newP := make([]int, len(path))
                copy(newP, path)
                out = append(out, newP)
            }
            path = path[:len(path)-1]
            return
        }
        dfs(r.Left, sum, path)
        dfs(r.Right, sum, path)
        path = path[:len(path)-1]
    }
    dfs(root, 0, nil)
    return out
}