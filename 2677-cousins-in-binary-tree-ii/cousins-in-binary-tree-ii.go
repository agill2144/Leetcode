/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func replaceValueInTree(root *TreeNode) *TreeNode {
    levelSum := map[int]int{}
    var dfsLevelSum func(r *TreeNode, level int)
    dfsLevelSum = func(r *TreeNode, level int) {
        // base
        if r == nil {return}
        // logic
        levelSum[level]+=r.Val
        dfsLevelSum(r.Left, level+1)
        dfsLevelSum(r.Right, level+1)
    }
    dfsLevelSum(root, 0)

    var dfs func(r *TreeNode,level int)
    dfs = func(r *TreeNode, level int) {
        // base
        if r == nil {return}

        // logic
        childSum := 0
        if r.Left != nil {childSum += r.Left.Val}
        if r.Right != nil {childSum += r.Right.Val}
        nextLevelSum := levelSum[level+1]
        if r.Left != nil {r.Left.Val = nextLevelSum - childSum}
        if r.Right != nil {r.Right.Val = nextLevelSum - childSum}

        dfs(r.Left, level+1)
        dfs(r.Right, level+1)
    }
    dfs(root, 0)
    root.Val = 0
    return root
}