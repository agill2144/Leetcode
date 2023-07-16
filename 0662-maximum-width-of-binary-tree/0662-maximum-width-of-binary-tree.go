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
            /*
                Once we know the first and last idx position of a level, then
                the width math is the same as sliding window size ( right-left+1 )
                or lastIdx-firstIdx+1
                
                But we have to consider null nodes in between too....
                instead of adding null nodes and using ptrs and vars to keep track of queue state
                we can enqueue a pair ( node, idx ), initially the idx will be 0 because starting node
                
                
                if we know the index of a node is C, then
                we can define the index of its left child node as 2*c ( if there is a left child )
                and for its right child it will be 2*c+1 ( if there is a right child )
               
                
            */
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