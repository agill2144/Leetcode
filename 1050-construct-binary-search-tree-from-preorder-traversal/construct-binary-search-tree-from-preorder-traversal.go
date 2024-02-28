/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
    inorder := make([]int, len(preorder))
    copy(inorder, preorder)
    sort.Ints(inorder)
    inorderIdx := map[int]int{}
    for i := 0; i < len(inorder); i++ {
        inorderIdx[inorder[i]] = i
    }
    
    prePtr := 0
    var dfs func(left, right int) *TreeNode
    dfs = func(left, right int) *TreeNode {
        // base
        if left > right {return nil}

        // logic
        rootVal := preorder[prePtr]
        prePtr++
        root := &TreeNode{Val: rootVal}
        rootIdx := inorderIdx[rootVal]
        root.Left = dfs(left, rootIdx-1)
        root.Right = dfs(rootIdx+1, right)
        return root
    }
    return dfs(0, len(inorder)-1)
    
}