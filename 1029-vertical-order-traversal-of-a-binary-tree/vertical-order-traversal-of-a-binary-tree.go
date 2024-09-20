/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// nodes at the same level AND AT THE SAME COL/WIDTH need to be sorted
func verticalTraversal(root *TreeNode) [][]int {
    if root == nil {return nil}
    colToNodes := map[int][]int{}
    minCol := math.MaxInt64
    maxCol := math.MinInt64
    type qNode struct {
        node *TreeNode
        col int // width
    }
    q := []*qNode{&qNode{root,0}}
    for len(q) != 0 {
        levelColToNodes := map[int][]int{}
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            cn := dq.node
            cc := dq.col
            levelColToNodes[cc] = append(levelColToNodes[cc], cn.Val)
            if cn.Left != nil {
                q = append(q, &qNode{cn.Left, cc-1})
            }
            if cn.Right != nil {
                q = append(q, &qNode{cn.Right, cc+1})
            }
            minCol=min(minCol, cc)
            maxCol=max(maxCol, cc)            
            qSize--
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