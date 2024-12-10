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
        col int
    }
    minCol := math.MaxInt64
    maxCol := math.MinInt64
    q := []*qNode{&qNode{root,0}}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        currNode := dq.node
        currCol := dq.col
        minCol = min(minCol, currCol)
        maxCol = max(maxCol, currCol)
        if colToNodes[currCol] == nil {colToNodes[currCol] = []int{}}
        colToNodes[currCol] = append(colToNodes[currCol], currNode.Val)
        if currNode.Left != nil {
            q = append(q, &qNode{currNode.Left, currCol-1})
        }
        if currNode.Right != nil {
            q = append(q, &qNode{currNode.Right, currCol+1})
        }
    }
    out := [][]int{}
    for i := minCol; i <= maxCol; i++ {
        out = append(out, colToNodes[i])
    }
    return out
}