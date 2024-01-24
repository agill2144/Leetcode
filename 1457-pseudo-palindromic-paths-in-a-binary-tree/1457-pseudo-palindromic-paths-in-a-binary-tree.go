/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func pseudoPalindromicPaths (root *TreeNode) int {
    count := 0
    var dfs func(r *TreeNode, path map[int]int)
    dfs = func(r *TreeNode, path map[int]int) {
        // base
        if r == nil {return}
        
        // logic
        path[r.Val]++
        dfs(r.Left, path)
        dfs(r.Right, path)
        if r.Left == nil && r.Right == nil && isPalindrome(path) {
            count++
        }
        path[r.Val]--
        if val := path[r.Val]; val == 0 {delete(path, r.Val)}
    }
    dfs(root, map[int]int{})
    return count
}
    
func isPalindrome( path map[int]int) bool {
    countOddFreqs := 0
    for _, v := range path {
        if v % 2 != 0 {
            countOddFreqs++
        }
    }
    return countOddFreqs <= 1
}