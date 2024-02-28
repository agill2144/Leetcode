/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 /*
    why we cant use dfs while maintain nodes in a col ( i.e width ) ?
    - draw it out with 4 levels of tree and see the problem
    - the order matters, top nodes in a col should be placed first
    - with pre-order dfs, top nodes will not always be placed first
    - therefore inorder to respect top nodes to be placed first, and then bottom nodes
    - this is why we need bfs
    - level order, top nodes are processed first and then bottom nodes
    - to group together nodes per col, we can save them in a map { $colNum: [nodes..] }
     
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