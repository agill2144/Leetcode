/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type qNode struct {
    node *TreeNode
    col int
}

func verticalOrder(root *TreeNode) [][]int {
    if root == nil {return nil}
    colToNodes := map[int][]int{}
    minCol := math.MaxInt64
    maxCol := math.MinInt64
    q := []*qNode{&qNode{root,0}}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        minCol = min(minCol, dq.col)
        maxCol = max(maxCol, dq.col)
        colToNodes[dq.col] = append(colToNodes[dq.col], dq.node.Val)
        if dq.node.Left != nil {
            q = append(q, &qNode{dq.node.Left, dq.col-1})
        }
        if dq.node.Right != nil {
            q = append(q, &qNode{dq.node.Right, dq.col+1})
        }
    }
    out := [][]int{}
    for i := minCol; i <= maxCol; i++ {
        out = append(out, colToNodes[i])
    }
    return out
}