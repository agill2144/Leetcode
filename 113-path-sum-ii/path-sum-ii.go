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
        path = append(path, r.Val)
        sum += r.Val
        if r.Left == nil && r.Right == nil {
            if sum == target {
                newL := make([]int, len(path))
                copy(newL, path)
                out = append(out, newL)
            }
            return
        } 
        dfs(r.Left, sum, path)
        dfs(r.Right, sum, path)   
        sum -= r.Val            
        path = path[:len(path)-1]
    }
    dfs(root, 0, nil)
    return out
}