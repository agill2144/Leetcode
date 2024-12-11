/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // we need width of each level = bfs 
 // idx tree nodes using heap idxs
 // left child = 2*i
 // right child = 2*i+1
 // start i = 1
func widthOfBinaryTree(root *TreeNode) int {
    idx := map[int][]int{} // { $level: [minIdx, maxIdx]}
    maxLevel := 0
    var dfs func(r *TreeNode, currIdx int, level int)
    dfs = func(r *TreeNode, currIdx, level int) {
        // base
        if r == nil {return}

        // logic
        maxLevel = max(maxLevel, level)
        if _, ok := idx[level]; !ok {idx[level] = []int{math.MaxInt64, math.MinInt64}}
        idx[level][0] = min(idx[level][0], currIdx)
        idx[level][1] = max(idx[level][1], currIdx)
        dfs(r.Left, 2*currIdx+1, level+1)
        dfs(r.Right, 2*currIdx+2, level+1)
    }
    dfs(root, 0, 0)
    maxW := -1
    for i := maxLevel; i >= 0; i-- {
        startIdx, endIdx := idx[i][0], idx[i][1]
        maxW = max(maxW, endIdx-startIdx+1)
    }
    return maxW
}