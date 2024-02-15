/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
    Same as https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
    - but start from the back of postorder instead
    - postorder = left-right-root
    - anytime we build trees, we start from root
    - root is the last element in postorder
    - and walk back , use inorder to find its right and left subtree

    - why are we building right subtree first instead of left subtree?
    - becuase of the order we are given in
    - we are given postorder, and postorder flows in; left,right,root
    - and if we are start from the back of postorder, we are starting from root
    - and walking back from root, it will be right, therefore we will be seeing right subtree elements first
    - in pre-order, we built left first and then right because preorder list is generated using preorder flow ; root,left,right
        - i.e root first, and then left, therefore we have left subtree elements first, therefore build left first
        - postorder is opposite of this
*/
func buildTree(inorder []int, postorder []int) *TreeNode {
    idxMap := map[int]int{}
    for i := 0; i < len(inorder); i++ {
        idxMap[inorder[i]] = i
    }
    post := len(postorder)-1
    var dfs func(left, right int) *TreeNode
    dfs = func(left, right int) *TreeNode {
        // base
        if left > right {return nil}

        // logic
        rootVal := postorder[post]
        post--
        root := &TreeNode{Val: rootVal}
        rootIdx := idxMap[rootVal]
        root.Right = dfs(rootIdx+1, right)
        root.Left = dfs(left, rootIdx-1)
        return root
    }
    return dfs(0, len(inorder)-1)
}