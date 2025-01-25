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
    colToNodes := map[int][]int{}
    minCol := math.MaxInt64
    maxCol := math.MinInt64
    out := [][]int{}
    type qNode struct {
        node *TreeNode
        col int
    }
    q := []*qNode{&qNode{root,0}}
    for len(q) != 0 {
        levelColToNodes := map[int][]int{}
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            qSize--
            node := dq.node
            col := dq.col
            minCol = min(minCol, col)
            maxCol = max(maxCol, col)
            levelColToNodes[col] = append(levelColToNodes[col], node.Val)
            if node.Left != nil {
                q = append(q, &qNode{node.Left,col-1})
            }
            if node.Right != nil {
                q = append(q, &qNode{node.Right, col+1})
            }
        }
        for col, nodes := range levelColToNodes {
            sort.Ints(nodes)
            colToNodes[col] = append(colToNodes[col], nodes...)
        }
    }
    for i := minCol; i <= maxCol; i++ {
        out = append(out, colToNodes[i])
    }
    return out
}
