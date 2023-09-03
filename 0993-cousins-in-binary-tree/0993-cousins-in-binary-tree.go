/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
    var xParent *TreeNode
    var yParent *TreeNode
    q := [][]*TreeNode{{root, nil}}    
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            node := dq[0]
            parent := dq[1]

            if node.Val == x {
                xParent = parent
            }
            if node.Val == y {
                yParent = parent
            }
            
            if node.Left != nil {
                q = append(q, []*TreeNode{node.Left, node})
            }
            if node.Right != nil {
                q = append(q, []*TreeNode{node.Right, node})
            }
            qSize--
        }
        if xParent == nil && yParent == nil {continue}
        if (xParent == nil && yParent != nil) || (xParent != nil && yParent == nil) {return false} 
        return xParent != yParent
    }
    return false
}