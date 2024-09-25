/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(postorder) == 0 || len(inorder) == 0 {return nil}
    post := len(postorder)-1
    inorderIdx := map[int]int{}
    for i := 0; i < len(inorder); i++ {inorderIdx[inorder[i]] = i}
    var dfs func(left, right int) *TreeNode
    dfs = func(left, right int) *TreeNode{
        // base
        if left > right {return nil}

        // logic
        // post = left -> right -> root

        // first we have the root node
        root := &TreeNode{Val: postorder[post]}
        post--
        idx := inorderIdx[root.Val]
        // after root, its right child ( look at post order seq )
        root.Right = dfs(idx+1, right)
        // after right, its left child ( look at post order seq )
        root.Left = dfs(left, idx-1)
        return root
    }
    return dfs(0, len(inorder)-1)

}