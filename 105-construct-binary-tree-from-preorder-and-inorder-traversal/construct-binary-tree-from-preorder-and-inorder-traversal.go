/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    pre := 0
    idx := map[int]int{}
    for i := 0; i < len(inorder); i++ {
        idx[inorder[i]] = i
    }
    var dfs func(left, right int) *TreeNode
    dfs = func(left, right int) *TreeNode {
        // base
        if left > right {return nil}

        // logic
        root := &TreeNode{Val: preorder[pre]}
        pre++
        rootIdx := idx[root.Val]
        root.Left = dfs(left, rootIdx-1)
        root.Right = dfs(rootIdx+1, right)
        return root
    }
    return dfs(0, len(preorder)-1)
}