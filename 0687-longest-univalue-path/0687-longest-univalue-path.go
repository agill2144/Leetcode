/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// dont understand AT ALL BUT IMPORTANT enough to get bookmarked
func longestUnivaluePath(root *TreeNode) int {
    maxLen := 0
    var dfs func(r *TreeNode) (int)
    dfs = func(r *TreeNode) (int) {
        // base
        if r == nil {return 0}
        
        // logic
        left := dfs(r.Left)
        right := dfs(r.Right)
        aLeft, aRight := 0,0
        
        if r.Left != nil && r.Val == r.Left.Val {
            aLeft += (left+1)
        }
        if r.Right != nil && r.Val == r.Right.Val {
            aRight += (right+1)
        }
        maxLen = max(maxLen, aLeft+aRight)
        return max(aLeft, aRight)
    }
    dfs(root)
    return maxLen
}

func max(x, y int) int {
    if x > y {return x}
    return y
}