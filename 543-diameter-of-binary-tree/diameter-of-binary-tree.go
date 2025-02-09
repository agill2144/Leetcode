/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {return 0}
    maxDia := 0
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}

        // logic
        leftDepth := getMaxDepth(r.Left)
        rightDepth := getMaxDepth(r.Right)
        maxDia = max(maxDia, leftDepth+rightDepth)
        dfs(r.Left)
        dfs(r.Right)

    }
    dfs(root)
    return maxDia
}


func getMaxDepth(r *TreeNode) int {
    maxDepth := 0
    var dfs func(r *TreeNode, depth int) 
    dfs = func(r *TreeNode, depth int) {
        // base
        maxDepth = max(maxDepth, depth)
        if r == nil {return}

        // logic
        dfs(r.Left, depth+1)
        dfs(r.Right, depth+1)
    }
    dfs(r, 0)
    return maxDepth
}