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
        // add node to path
        path = append(path, r.Val)
        // add node value to sum
        sum += r.Val

        // we have updated sum, now
        // check whether this node is a leaf
        // if yes, and the sum cumulated so far == target
        // make a copy of the path, save the copy to output
        if r.Left == nil && r.Right == nil {
            if sum == targetSum {
                newP := make([]int, len(path))
                copy(newP, path)
                out = append(out, newP)
            }
        }

        dfs(r.Left, sum, path)
        dfs(r.Right, sum, path)
        path = path[:len(path)-1]
    }
    dfs(root, 0, nil)
    return out
}