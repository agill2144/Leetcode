/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
    approach : bfs
    - maintain cell number / width for each node
    - store each node to their cell list in map : cellToNodes
*/
func verticalOrder(root *TreeNode) [][]int {
    if root == nil {return nil}
    widthToNodes := map[int][]int{}
    min := math.MaxInt64
    max := math.MinInt64

    q := [][]interface{}{{root, 0}}
    for len(q) != 0 {
        
        dq := q[0]
        q = q[1:]
        currNode := dq[0].(*TreeNode)
        currWidth := dq[1].(int)
        
        if currWidth < min {min = currWidth}
        if currWidth > max {max = currWidth}

        widthToNodes[currWidth] = append(widthToNodes[currWidth], currNode.Val)
        if currNode.Left != nil {
            q = append(q, []interface{}{currNode.Left, currWidth-1})
        }
        if currNode.Right != nil {
            q = append(q, []interface{}{currNode.Right, currWidth+1})
        }
        
    }
    out := [][]int{}
    for i := min; i <= max; i++ {
        out = append(out, widthToNodes[i])
    }
    return out
}