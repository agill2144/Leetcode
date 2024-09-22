/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func widthOfBinaryTree(root *TreeNode) int {
    type qNode struct {
        node *TreeNode
        idx int
    }
    q := []*qNode{&qNode{root,1}}
    maxW := -1
    for len(q) != 0 {
        qSize := len(q)
        start, end := q[0].idx, q[qSize-1].idx
        maxW = max(maxW, end-start+1)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            cn := dq.node
            if cn.Left != nil {
                q = append(q, &qNode{cn.Left, 2*dq.idx})
            }
            if cn.Right != nil {
                q = append(q, &qNode{cn.Right, 2*dq.idx+1})
            }
            qSize--
        }
    }
    return maxW
}