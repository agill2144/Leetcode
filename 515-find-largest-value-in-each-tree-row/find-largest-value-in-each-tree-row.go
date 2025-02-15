/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
    if root == nil {return nil}
    levelMax := map[int]int{}
    maxLevel := 0
    var dfs func(r *TreeNode, level int)
    dfs = func(r *TreeNode, level int) {
        // base
        if r == nil {return}

        // logic
        maxLevel = max(maxLevel, level)
        val, ok := levelMax[level]
        if ok {
            levelMax[level] = max(r.Val, val)
        } else if !ok {
            levelMax[level] = r.Val
        }
        dfs(r.Left,level+1)
        dfs(r.Right, level+1)
    }
    dfs(root, 0)
    res := []int{}
    for i := 0; i <= maxLevel; i++ {
        res = append(res, levelMax[i])
    }
    return res
}