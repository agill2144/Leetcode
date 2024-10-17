/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func verticalTraversal(root *TreeNode) [][]int {
    widthToNodes := map[int][]int{}
    type qNode struct {
        node *TreeNode
        row int
        col int
    }
    minC := math.MaxInt64
    maxC := math.MinInt64
    q := []*qNode{&qNode{root,0,0}}
    for len(q) != 0 {
        qSize := len(q)
        levelWidthToNodes := map[int][]int{}
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            node := dq.node
            row := dq.row
            col := dq.col
            minC = min(minC, col)
            maxC = max(maxC, col)
            levelWidthToNodes[col] = append(levelWidthToNodes[col], node.Val)
            if node.Left != nil {
                q = append(q, &qNode{node.Left, row+1,col-1})
            }
            if node.Right != nil {
                q = append(q, &qNode{node.Right, row+1,col+1})
            }
            qSize--
        }
        for k, v := range levelWidthToNodes {
            sort.Ints(v)
            widthToNodes[k] = append(widthToNodes[k], v...)
        }
    }
    out := [][]int{}
    for i := minC; i <= maxC; i++ {
        out = append(out, widthToNodes[i])
    }
    return out
}