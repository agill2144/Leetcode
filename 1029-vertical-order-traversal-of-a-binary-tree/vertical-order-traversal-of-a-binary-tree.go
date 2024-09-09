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
        width int
    }
    tmp := map[int][]int{}
    level := 0
    q := []*qNode{&qNode{root,0}}
    minW := math.MaxInt64
    maxW := math.MinInt64
    for len(q) != 0 {
        qSize := len(q)
        levelItems := map[int][]int{}
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            cn := dq.node
            cw := dq.width
            minW = min(minW, cw)
            maxW = max(maxW, cw)
            if tmp[cw] == nil {tmp[cw] = []int{}}
            levelItems[cw] = append(levelItems[cw], cn.Val)
            if cn.Left != nil {
                q = append(q, &qNode{cn.Left, cw-1})
            }
            if cn.Right != nil {
                q = append(q, &qNode{cn.Right, cw+1})
            }
            qSize--
        }
        for w, items := range levelItems{
            if tmp[w] == nil {tmp[w] = []int{}}
            sort.Ints(items)
            tmp[w] = append(tmp[w], items...)
        }
        level++
    }
    out := [][]int{}
    for i := minW; i <= maxW; i++ {
        out = append(out, tmp[i])
    }
    return out
}