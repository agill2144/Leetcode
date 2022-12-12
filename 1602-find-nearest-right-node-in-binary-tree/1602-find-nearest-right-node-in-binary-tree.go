/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findNearestRightNode(root *TreeNode, u *TreeNode) *TreeNode {
    
    q := []*TreeNode{root}
    for len(q) != 0 {
        
        qSize := len(q)
        for qSize != 0 {

            dq := q[0]
            q = q[1:]
            qSize--

            if dq == u {
                if qSize > 0 {return q[0]}
            }
            if dq.Left != nil {
                q = append(q, dq.Left)
            }
            
            if dq.Right != nil {
                q = append(q, dq.Right)            
            }
        }
        
    }
    return nil
    
}