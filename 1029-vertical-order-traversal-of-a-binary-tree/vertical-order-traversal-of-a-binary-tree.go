/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func verticalTraversal(root *TreeNode) [][]int {
    type qNode struct {
        node *TreeNode
        col int
    }
    colToNodes := map[int][]int{}
    q := []*qNode{&qNode{root,0}}
    minCol := math.MaxInt64
    maxCol := math.MinInt64
    for len(q) != 0 {
        qSize := len(q) 
        levelColToNodes := map[int][]int{}
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            qSize--
            currNode := dq.node
            currCol := dq.col
            minCol = min(minCol, currCol)
            maxCol = max(maxCol, currCol)
            levelColToNodes[currCol] = append(levelColToNodes[currCol], currNode.Val)
            if currNode.Left != nil {
                q = append(q, &qNode{currNode.Left,currCol-1})
            } 
            if currNode.Right != nil {
                q = append(q, &qNode{currNode.Right,currCol+1})
            } 
        }
        for level, nodes := range levelColToNodes {
            if len(nodes) > 1 {sort.Ints(nodes)}
            colToNodes[level] = append(colToNodes[level], nodes...)
        }
    }
    out := [][]int{}
    for i := minCol; i <= maxCol; i++ {
        out = append(out, colToNodes[i])
    }
    return out
}