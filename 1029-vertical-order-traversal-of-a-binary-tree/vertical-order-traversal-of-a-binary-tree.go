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
        // nodes on a specific level at the same width, need to be sorted by their values
        // we have level specific nodes in the tmp map
        // we need to sort if nodes are on the same width
        // nodes in tmp map are stored by each node's width
        // therefore we can loop over the tmp map ( which is level specific withToNodes map )
        // and sort the list of nodes
        for w, items := range tmp {
            sort.Ints(items)
            // then we can push the sorted width of nodes from this level to global widthToNodes map
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