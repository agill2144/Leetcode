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
    type qNode struct {
        node *TreeNode
        col int
    }
    minCol := math.MaxInt64
    maxCol := math.MinInt64
    q := []*qNode{&qNode{root,0}}
    for len(q) != 0 {
        qSize := len(q)
        levelColToNodes := map[int][]int{}
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            qSize--
            cn := dq.node
            cc := dq.col
            minCol = min(minCol, cc)
            maxCol = max(maxCol, cc)
            levelColToNodes[cc] = append(levelColToNodes[cc], cn.Val)
            if cn.Left != nil {
                q = append(q, &qNode{cn.Left, cc-1})
            }
            if cn.Right != nil {
                q = append(q, &qNode{cn.Right, cc+1})
            }
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