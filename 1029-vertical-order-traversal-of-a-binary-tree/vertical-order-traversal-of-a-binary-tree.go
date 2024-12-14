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
// func verticalTraversal(root *TreeNode) [][]int {
//     if root == nil {return nil}
//     type qNode struct {
//         node *TreeNode
//         col int
//     }
//     minCol := math.MaxInt64
//     maxCol := math.MinInt64
//     colToNodes := map[int][]int{}
//     q := []*qNode{&qNode{root,0}}
//     for len(q) != 0 {
//         qSize := len(q)
//         levelColToNodes := map[int][]int{}
//         for qSize != 0 {
//             dq := q[0]
//             q = q[1:]
//             qSize--
//             curr := dq.node
//             col := dq.col
//             minCol = min(minCol, col)
//             maxCol = max(maxCol, col)
//             levelColToNodes[col] = append(levelColToNodes[col], curr.Val)
//             if curr.Left != nil {
//                 q = append(q, &qNode{curr.Left, col-1})
//             }
//             if curr.Right != nil {
//                 q = append(q, &qNode{curr.Right, col+1})
//             }
//         }
//         for col, nodes := range levelColToNodes {
//             sort.Ints(nodes)
//             colToNodes[col] = append(colToNodes[col], nodes...)
//         }
//     }
//     out := [][]int{}
//     for i := minCol; i <= maxCol; i++ {
//         out = append(out, colToNodes[i])
//     }
//     return out
// }