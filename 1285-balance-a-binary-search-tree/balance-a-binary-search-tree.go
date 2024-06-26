/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func balanceBST(root *TreeNode) *TreeNode {
    sortedArr := []int{}
    var inorder func(r *TreeNode)
    inorder = func(r *TreeNode) {
        // base
        if r == nil {return}

        // logic
        inorder(r.Left)
        sortedArr = append(sortedArr, r.Val)
        inorder(r.Right)
    }
    inorder(root)

    var dfs func(left, right int) *TreeNode
    dfs = func(left, right int) *TreeNode {
        // base
        if left > right {return nil}

        // logic
        mid := left + (right-left)/2
        root := &TreeNode{Val: sortedArr[mid]}
        root.Left = dfs(left, mid-1)
        root.Right = dfs(mid+1, right)
        return root
    }
    return dfs(0, len(sortedArr)-1)
}