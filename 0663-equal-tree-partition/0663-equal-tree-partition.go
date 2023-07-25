/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// remember that a number can be ONLY halved if its even number
// ex; 36 can be halved ( 13 )
// 37 cannot be because its impossible to have a perfect whole half
// thats the core logic
// Need a HashMap to handle the case when total sum is 0. If totalSum == 0, need to check if the count of 0 is larger than 
func checkEqualTree(root *TreeNode) bool {
    
    sums := map[int]int{}
    var dfs func(r *TreeNode) int
    dfs = func(r *TreeNode) int {
        // base
        if r == nil {return 0}
        
        // logic
        left := dfs(r.Left)
        right := dfs(r.Right)
        curr := left+right+r.Val
        sums[curr]++
        return curr
    }
    total := dfs(root)
    if total % 2 != 0 {return false}
    
    c, ok := sums[total/2]
    if total == 0 {return c > 1}
    return ok
}