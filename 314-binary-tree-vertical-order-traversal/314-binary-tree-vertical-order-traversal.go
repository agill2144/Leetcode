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
    type pair struct {
        node *TreeNode
        width int
    }
    wMap := map[int][]int{}
    q := []*pair{&pair{node:root,width:0}}
    min := math.MaxInt64
    max := math.MinInt64
    for len(q) != 0 {
        dqPair := q[0]
        q = q[1:]
        currNode := dqPair.node
        currWidth := dqPair.width
        wMap[currWidth] = append(wMap[currWidth], currNode.Val)
        if currWidth < min {min = currWidth}
        if currWidth > max {max = currWidth}
        
        if currNode.Left != nil {
            q = append(q, &pair{node:currNode.Left, width: currWidth-1})
        }
        if currNode.Right != nil {
            q = append(q, &pair{node:currNode.Right, width: currWidth+1})
        }
    }
    
    result := [][]int{}
    for i := min; i <= max; i++ {
        result = append(result, wMap[i])
    }
    return result
}