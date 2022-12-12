/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) != len(inorder){return nil}
    
    idxMap := map[int]int{}
    for i := 0; i < len(inorder); i++ {idxMap[inorder[i]] = i}
    prePtr := 0
    var dfs func(left, right int) *TreeNode
    dfs = func(left, right int) *TreeNode{
        // base
        if left > right {return nil}
       
        // logic
        rootVal := preorder[prePtr]
        prePtr++
        rootIdx := idxMap[rootVal]
        root := &TreeNode{Val: rootVal}
        root.Left = dfs(left, rootIdx-1)
        root.Right = dfs(rootIdx+1, right)
        return root
    }
    
    return dfs(0, len(preorder)-1)
    
}