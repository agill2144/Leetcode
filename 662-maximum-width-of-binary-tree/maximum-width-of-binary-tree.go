/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // we need width of each level = bfs 
 // idx tree nodes using heap idxs
 // left child = 2*i
 // right child = 2*i+1
 // start i = 1
func widthOfBinaryTree(root *TreeNode) int {
    if root == nil {return 0}
    type qNode struct {
        idx int
        node *TreeNode
    }   
    q := []*qNode{&qNode{1,root}}
    maxW := 0
    for len(q) != 0 {
        qSize := len(q)
        startIdx := q[0].idx
        endIdx := q[qSize-1].idx
        maxW = max(maxW, endIdx-startIdx+1)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if dq.node.Left != nil {
                q = append(q, &qNode{2*dq.idx, dq.node.Left})
            }
            if dq.node.Right != nil {
                q = append(q, &qNode{2*dq.idx+1, dq.node.Right})
            }
            qSize--
        }
    }
    return maxW
}