/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func str2tree(s string) *TreeNode {
    i := 0
    var dfs func() *TreeNode
    dfs = func() *TreeNode {
        // base
        if i == len(s) {return nil}
        
        // logic
        num := 0
        isNeg := false
        if s[i] == '-' {isNeg = true; i++}
        for i < len(s) && s[i] >= '0' && s[i] <= '9'  {
            num = num * 10 + int(s[i]-'0')
            i++
        }
        if isNeg {num = num*-1}
        root := &TreeNode{Val: num}
        if i < len(s) && s[i] == '(' {
            i++
            root.Left = dfs()
        }
        if i < len(s) && s[i] == '(' && root.Left != nil {
            i++
            root.Right = dfs()
        }
        i++
        return root
    }
    return dfs()
}