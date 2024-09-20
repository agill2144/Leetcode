/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func verticalOrder(root *TreeNode) [][]int {
    if root == nil {return nil}
    colToNodes := map[int][]int{}
    type qNode struct {
        node *TreeNode
        row int // level
        col int // width
    }
    minCol := math.MaxInt64
    maxCol := math.MinInt64
    q := []*qNode{&qNode{root,0,0}}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        cn := dq.node
        cr := dq.row
        cc := dq.col
        minCol = min(cc, minCol)
        maxCol = max(cc, maxCol)
        colToNodes[cc] = append(colToNodes[cc], cn.Val)
        if cn.Left != nil {
            q = append(q, &qNode{cn.Left, cr+1, cc-1})
        }
        if cn.Right != nil {
            q = append(q, &qNode{cn.Right, cr+1, cc+1})
        }
    }

    out := [][]int{}
    for i := minCol; i <= maxCol; i++ {
        out = append(out, colToNodes[i])
    }
    return out

}