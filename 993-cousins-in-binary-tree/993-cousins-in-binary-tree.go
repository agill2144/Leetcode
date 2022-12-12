/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
    
    type pair struct {
        node *TreeNode
        parent *TreeNode
    }
    q := []*pair{}
    q = append(q, &pair{node: root, parent: nil})
    var xParent *TreeNode
    var yParent *TreeNode
    
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            dqNode := dq.node
            dqParent := dq.parent
            
            if dqNode.Val == x {xParent = dqParent}
            if dqNode.Val == y {yParent = dqParent}
            
            if dqNode.Left != nil {
                q = append(q, &pair{node: dqNode.Left, parent: dqNode})
            }
            if dqNode.Right != nil {
                q = append(q, &pair{node: dqNode.Right, parent: dqNode})
            }
            qSize--
        }
        
        if xParent != nil && yParent != nil {
            return xParent != yParent
        }
        if (xParent != nil && yParent == nil) || (xParent == nil && yParent != nil) {
            return false
        } 
    }
    
    return false
}