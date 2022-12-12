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
        if left > right {return nil}
        
        // logic
        rootVal := postorder[postRootPtr]
        postRootPtr--
        // find the root in inorder
        // this will tell us that the left side of the same idx in postorder is the left postorder tree
        // the discovering of root in inorder will also tell us the right subtree in postorder array
        rootIdx := idxMap[rootVal]
        root := &TreeNode{Val: rootVal}

        // why +1 and -1 : because we rootIdx is already used - it may not be the SAME VALUE in postorder as with inorder but thats the point, that idx is already use
        root.Right = dfs(rootIdx+1, right)
        root.Left = dfs(left, rootIdx-1)
        return root
    }
    return dfs(0, len(postorder)-1)
}