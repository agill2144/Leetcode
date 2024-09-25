/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {return nil}
    pre := 0
    inorderIdx := map[int]int{}
    for i := 0; i < len(inorder); i++ {inorderIdx[inorder[i]] = i}
    var dfs func(left, right int) *TreeNode
    dfs = func(left, right int) *TreeNode{
        // base
        if left > right {return nil}

        // logic
        root := &TreeNode{Val: preorder[pre]}
        pre++
        idx := inorderIdx[root.Val]
        root.Left = dfs(left, idx-1)
        root.Right = dfs(idx+1, right)
        return root
    }
    return dfs(0, len(inorder)-1)
    
}