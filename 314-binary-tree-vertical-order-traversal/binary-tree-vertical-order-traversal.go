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
    cols := map[int][]int{}
    minCol := math.MaxInt64
    maxCol := math.MinInt64
    type qNode struct {
        node *TreeNode
        col int
    }
    q := []*qNode{&qNode{root,0}}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        cn := dq.node
        cc := dq.col
        cols[cc] = append(cols[cc], cn.Val)
        minCol = min(minCol, cc)
        maxCol = max(maxCol, cc)
        if cn.Left != nil {
            q = append(q, &qNode{cn.Left, cc-1})
        }
        if cn.Right != nil {
            q = append(q, &qNode{cn.Right, cc+1})
        }
    }
    out := [][]int{}
    for i := minCol; i <= maxCol; i++ {
        out = append(out, cols[i])
    }
    return out
}