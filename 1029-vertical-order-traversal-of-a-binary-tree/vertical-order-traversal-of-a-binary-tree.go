/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func verticalTraversal(root *TreeNode) [][]int {
    if root == nil {return nil}
    type qNode struct {
        col int
        node *TreeNode
    }
    minCol := math.MaxInt64
    maxCol := math.MinInt64
    colToNodes := map[int][]int{}
    q := []*qNode{&qNode{0, root}}
    for len(q) != 0 {
        levelColToNodes := map[int][]int{}
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            node := dq.node
            col := dq.col
            minCol = min(minCol, col)
            maxCol = max(maxCol, col)
            levelColToNodes[col] = append(levelColToNodes[col], node.Val)
            if node.Left != nil {
                q = append(q, &qNode{col-1,node.Left})
            }
            if node.Right != nil {
                q = append(q, &qNode{col+1, node.Right})
            }
            qSize--
        }
        for col, nodes := range levelColToNodes {
            if len(nodes) > 1 {
                sort.Ints(nodes)
            }
            colToNodes[col] = append(colToNodes[col], nodes...)
        }
    }
    out := [][]int{}
    for i := minCol; i <= maxCol; i++ {
        out = append(out, colToNodes[i])
    }
    return out

}