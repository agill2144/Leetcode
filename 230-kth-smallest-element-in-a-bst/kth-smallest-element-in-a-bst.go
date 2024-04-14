/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // count to k while traversing inorder
 // inorder == sorted order
func kthSmallest(root *TreeNode, k int) int {
    count := 0
    var dfs func(r *TreeNode) (int, bool)
    dfs = func(r *TreeNode) (int, bool) {
        // base
        if r == nil {return -1, false}

        // logic
        left, lok := dfs(r.Left)
        if lok {return left, true}

        count++
        if count == k {
            return r.Val, true
        }

        right, rok := dfs(r.Right)
        if rok {return right, true}

        return -1, false
    }
    val, _ := dfs(root)
    return val
}