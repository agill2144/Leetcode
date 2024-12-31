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
            $colNum: {
                $levelNum: [node1, node2]
            }
        }
    */
    minCol := math.MaxInt64
    maxCol := math.MinInt64
    maxLevel := 0
    colToNodes := map[int]map[int][]int{}
    var dfs func(r *TreeNode, level int, col int)
    dfs = func(r *TreeNode, level int, col int) {
        // base
        if r == nil {return}
        // logic
        if colToNodes[col] == nil {colToNodes[col] = map[int][]int{}}
        if colToNodes[col][level] == nil {colToNodes[col][level] = []int{}}
        colToNodes[col][level] = append(colToNodes[col][level], r.Val)
        maxCol = max(maxCol, col)
        minCol = min(minCol, col)
        maxLevel = max(maxLevel, level)
        dfs(r.Left, level+1, col-1)
        dfs(r.Right, level+1, col+1)
    }
    dfs(root, 0, 0)
    out := [][]int{}
    for i := minCol; i<= maxCol; i++ {
        colData := []int{}
        for j := 0; j <= maxLevel; j++ {
            nodes := colToNodes[i][j]
            if len(nodes) == 0 {continue}
            sort.Ints(nodes)
            colData = append(colData, nodes...)
        }
        out = append(out, colData)
    }
    return out
}
// func verticalTraversal(root *TreeNode) [][]int {
//     if root == nil {return nil}
//     type qNode struct {
//         col int
//         node *TreeNode
//     }
//     minCol := math.MaxInt64
//     maxCol := math.MinInt64
//     colToNodes := map[int][]int{}
//     q := []*qNode{&qNode{0, root}}
//     for len(q) != 0 {
//         levelColToNodes := map[int][]int{}
//         qSize := len(q)
//         for qSize != 0 {
//             dq := q[0]
//             q = q[1:]
//             node := dq.node
//             col := dq.col
//             minCol = min(minCol, col)
//             maxCol = max(maxCol, col)
//             levelColToNodes[col] = append(levelColToNodes[col], node.Val)
//             if node.Left != nil {
//                 q = append(q, &qNode{col-1,node.Left})
//             }
//             if node.Right != nil {
//                 q = append(q, &qNode{col+1, node.Right})
//             }
//             qSize--
//         }
//         for col, nodes := range levelColToNodes {
//             if len(nodes) > 1 {
//                 sort.Ints(nodes)
//             }
//             colToNodes[col] = append(colToNodes[col], nodes...)
//         }
//     }
//     out := [][]int{}
//     for i := minCol; i <= maxCol; i++ {
//         out = append(out, colToNodes[i])
//     }
//     return out

// }