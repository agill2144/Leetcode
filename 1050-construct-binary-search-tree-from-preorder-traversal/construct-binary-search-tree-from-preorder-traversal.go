/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
    pre := 0
    var dfs func(mn, mx int) *TreeNode
    dfs = func(mn, mx int) *TreeNode {
        // base
        if pre == len(preorder) {return nil}
        if preorder[pre] <= mn || preorder[pre] >= mx {return nil}

        // logic
        root := &TreeNode{Val: preorder[pre]}
        pre++
        root.Left = dfs(mn, root.Val)
        root.Right = dfs(root.Val, mx)
        return root
    }
    return dfs(math.MinInt64, math.MaxInt64)

}