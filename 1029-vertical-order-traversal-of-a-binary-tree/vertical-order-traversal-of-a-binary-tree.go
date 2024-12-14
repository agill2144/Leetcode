/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func verticalTraversal(root *TreeNode) [][]int {
    /*
        nodes on the SAME ROW AND SAME COL need to be sorted
        we also want the ability to search by col because at the end
        we are storing from left to right side of the col
        the leftest col is min col and rightest col is max col val
        therefore store by cols as initial key.
        then we to group nodes that are on the same row within col
        hence using another map {rowNum: [nodes]}
        so that we can sort nodes in the same row and col; data[$col][$rowNum]  
        {
            $col: {$row: [nodes]}
        }
    */
    minCol := math.MaxInt64
    maxCol := math.MinInt64
    maxRow := math.MinInt64
    data := map[int]map[int][]int{}
    var dfs func(r *TreeNode, row, col int)
    dfs = func(r *TreeNode, row, col int) {
        // base
        if r == nil {return}
        // logic
        if data[col] == nil {data[col] = map[int][]int{}}
        data[col][row] = append(data[col][row], r.Val)
        minCol = min(minCol, col)
        maxCol = max(maxCol, col)
        maxRow = max(maxRow, row)
        dfs(r.Left, row+1, col-1)
        dfs(r.Right, row+1, col+1)
    }
    dfs(root, 0,0)
    out := [][]int{}
    for i := minCol; i <= maxCol; i++ {
        rows := data[i]
        colData := []int{}
        for j := 0; j <= maxRow; j++ {
            if rows[j] == nil {continue}
            if len(rows[j]) > 1 {sort.Ints(rows[j])}
            colData = append(colData, rows[j]...)
        }
        out = append(out, colData)
    }
    return out
}
