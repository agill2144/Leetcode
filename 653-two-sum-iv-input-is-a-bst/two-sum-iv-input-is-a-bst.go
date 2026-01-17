/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
    approach: inorder + two sum
    - inorder -> sorted array
    - then it becomes a sorted array and find two sum in a sorted array
        - binary search: o(nlogn)
        - two ptrs: o(n)
    - total tc: o(n) + o(n)
    - sc: dfs implicit stack: o(n) + o(n) for sorted array

    approach: dfs + compliment search ( order of dfs does not matter )
*/
func findTarget(root *TreeNode, k int) bool {
    seen := map[int]bool{}
    var dfs func(r *TreeNode) bool
    dfs = func(r *TreeNode) bool {
        // base
        if r == nil {return false}

        // logic
        if seen[k-r.Val] {return true}
        seen[r.Val] = true
        if dfs(r.Left) {return true}
        return dfs(r.Right)
    }
    return dfs(root)
}