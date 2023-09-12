/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func replaceValueInTree(root *TreeNode) *TreeNode {
    if root == nil {return root}
    levelSum := map[int]int{}
    var dfs1 func(r *TreeNode, level int)
    dfs1 = func(r *TreeNode, level int) {
        // base
        if r == nil {return}
        
        // logic
        levelSum[level] += r.Val
        dfs1(r.Left, level+1)
        dfs1(r.Right, level+1)
    }
    dfs1(root, 0)
    
    var dfs2 func(r *TreeNode, level int)
    dfs2 = func(r *TreeNode, level int) {
        // base
        if r == nil {return}
        
        // logic
        nextLevelSum, ok := levelSum[level+1]
        if !ok {return}
        siblingSum := 0
        if r.Left != nil {siblingSum += r.Left.Val}
        if r.Right != nil {siblingSum += r.Right.Val}
        cousinSum := nextLevelSum - siblingSum
        if r.Left != nil {r.Left.Val = cousinSum}
        if r.Right != nil {r.Right.Val = cousinSum}
        dfs2(r.Left, level+1)
        dfs2(r.Right, level+1)
    }
    dfs2(root, 0)
    root.Val = 0
    return root
}