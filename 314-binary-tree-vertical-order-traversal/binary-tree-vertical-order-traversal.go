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
    widthToNode := map[int][]int{}
    type qNode struct {
        node *TreeNode
        width int
    }
    startWidth := math.MaxInt64
    endWidth := math.MinInt64
    q := []*qNode{&qNode{root,0}}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        currNode := dq.node
        currWidth := dq.width
        startWidth = min(startWidth, currWidth)
        endWidth = max(endWidth, currWidth)
        widthToNode[currWidth] = append(widthToNode[currWidth], currNode.Val)
        if currNode.Left != nil {
            q = append(q, &qNode{currNode.Left,currWidth-1})
        }
        if currNode.Right != nil {
            q = append(q, &qNode{currNode.Right,currWidth+1})
        }
    }
    out := [][]int{}
    for i := startWidth; i <= endWidth; i++ {
        out = append(out, widthToNode[i])
    }
    return out
}