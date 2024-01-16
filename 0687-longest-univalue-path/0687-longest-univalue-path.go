/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {
    var dfs func(r *TreeNode) int
    ans := 0
    dfs = func(r *TreeNode) int {
        // base
        if r == nil {return 0}
        
        // logic
        left := dfs(r.Left)
        right := dfs(r.Right)
        

        // either left is long or right is long
        leftCount := 0
        rightCount := 0
        if r.Left != nil && r.Val == r.Left.Val { leftCount += left+1 }
        if r.Right != nil && r.Val == r.Right.Val { rightCount += right+1 }
        
        
        ans = max(ans, leftCount+rightCount)
        return max(leftCount, rightCount)              
    }
    
    dfs(root)
    return ans
}

func max(x, y int) int {
    if x > y {return x}
    return y
}