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
    var dfs func(r *TreeNode, sum int, path []int)
    dfs = func(r *TreeNode, sum int, path []int) {
        // base
        if r == nil {return}

        // logic
        path = append(path, r.Val)
        sum += r.Val
        dfs(r.Left, sum, path)
        dfs(r.Right, sum, path)
        if r.Left == nil && r.Right == nil {
            if sum == targetSum {
                newP := make([]int, len(path))
                copy(newP, path)
                out = append(out, newP)
            }
        }
        path = path[:len(path)-1]
    }
    dfs(root, 0, nil)
    return out
}