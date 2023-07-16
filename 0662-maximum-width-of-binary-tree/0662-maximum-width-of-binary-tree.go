/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func widthOfBinaryTree(root *TreeNode) int {
    if root == nil {return 0}
    type pair struct {
        node *TreeNode
        idx int
    }
    q := []*pair{&pair{node:root,idx:0}}
    max := 0
    for len(q) != 0 {
        qSize := len(q)
        headIdx := q[0].idx
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            node := dq.node
            if  node.Left != nil {
                q = append(q, &pair{node:node.Left, idx:2*dq.idx})
            }
            if  node.Right != nil {
                q = append(q, &pair{node:node.Right, idx:2*dq.idx+1})
            }
            qSize--
            
            // headIdx is for this level
            // we are processing nodes on this level
            if dq.idx - headIdx + 1 > max {
                max = dq.idx - headIdx+1
            }
            
        }
    }
    return max
}