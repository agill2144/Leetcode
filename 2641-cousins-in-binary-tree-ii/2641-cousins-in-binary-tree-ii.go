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
    dfs1 = func(r *TreeNode, level int){
        // base
        if r == nil {return}
        
        // logic
        levelSum[level]+=r.Val
        dfs1(r.Left, level+1)
        dfs1(r.Right, level+1)
    }
    dfs1(root, 0)
    
    var dfs func(r *TreeNode, level, siblingSum int)
    dfs = func(r *TreeNode, level , siblingSum int) {
        // base
        if r == nil {return}
        
        // logic
        r.Val = levelSum[level]-siblingSum
        nextSiblingSum := 0
        if r.Left != nil {nextSiblingSum += r.Left.Val}
        if r.Right != nil {nextSiblingSum += r.Right.Val}
        dfs(r.Left, level+1, nextSiblingSum)
        dfs(r.Right, level+1, nextSiblingSum)
    }
    dfs(root, 0, root.Val)
    return root
}