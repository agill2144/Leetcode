/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 /*
    approach: brute force
    - assuming the preorder is valid for a BST
    - we can copy that into a new array, sort it
    - this new sorted array is our inorder 
    - now we have 1. preorder and 2. inorder arrays
    - This question now becomes a traditional "construct tree using preorder and inorder"

    time = o(n) + o(nlogn) + o(n) = o(nlogn)
    space = o(n)
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