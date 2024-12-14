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
        node *TreeNode
        col int
    }
    minCol := math.MaxInt64
    maxCol := math.MinInt64
    colToNodes := map[int][]int{}
    q := []*qNode{&qNode{root,0}}
    for len(q) != 0 {
        qSize := len(q)
        levelColToNodes := map[int][]int{}
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            qSize--
            curr := dq.node
            col := dq.col
            minCol = min(minCol, col)
            maxCol = max(maxCol, col)
            levelColToNodes[col] = append(levelColToNodes[col], curr.Val)
            if curr.Left != nil {
                q = append(q, &qNode{curr.Left, col-1})
            }
            if curr.Right != nil {
                q = append(q, &qNode{curr.Right, col+1})
            }
        }
        for col, nodes := range levelColToNodes {
            sort.Ints(nodes)
            colToNodes[col] = append(colToNodes[col], nodes...)
        }
    }
    out := [][]int{}
    for i := minCol; i <= maxCol; i++ {
        out = append(out, colToNodes[i])
    }
    return out
}