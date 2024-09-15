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
    widthToNodes := map[int][]int{}
    type qNode struct {
        node *TreeNode
        width int
    }
    minWidth := math.MaxInt64
    maxWidth := math.MinInt64
    level := 0
    q := []*qNode{&qNode{root,0}}
    for len(q) != 0 {
        qSize := len(q)
        tmp := map[int][]int{} // width to nodes in this $level
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            cn := dq.node
            cw := dq.width
            tmp[cw] = append(tmp[cw], cn.Val)
            if cn.Left != nil {
                q = append(q, &qNode{cn.Left, cw-1})
            }
            if cn.Right != nil {
                q = append(q, &qNode{cn.Right, cw+1})
            }
            minWidth = min(minWidth, cw)
            maxWidth = max(maxWidth, cw)
            qSize--
        }
        for w, items := range tmp {
            sort.Ints(items)
            widthToNodes[w] = append(widthToNodes[w], items...)
        }
        level++
    }
    out := [][]int{}
    for i := minWidth; i <= maxWidth; i++ {
        out = append(out, widthToNodes[i])
    }
    return out
}