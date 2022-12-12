/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    idxMap := map[int]int{}
    for i := 0; i < len(inorder); i++ {
        idxMap[inorder[i]] = i
    }
    postRootPtr := len(postorder)-1
    
    var dfs func(left, right int) *TreeNode
    dfs = func(left, right int) *TreeNode {
        // base
        if left > right {
            return nil
        }
        // logic
        rootVal := postorder[postRootPtr]
        postRootPtr--
        rootIdx := idxMap[rootVal]
        root := &TreeNode{Val: rootVal}
        root.Right = dfs(rootIdx+1,right)
        root.Left = dfs(left, rootIdx-1)
        return root
    }
    return dfs(0, len(postorder)-1)
}