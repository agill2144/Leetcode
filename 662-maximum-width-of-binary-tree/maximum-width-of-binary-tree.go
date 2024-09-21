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
    type qNode struct {
        node *TreeNode
        idx int
    }
    maxWidth := 1
    q := []*qNode{&qNode{root, 1}}
    for len(q) != 0 {
        qSize := len(q)

        firstIdx := q[0].idx
        lastIdx := q[qSize-1].idx
        maxWidth = max(maxWidth, lastIdx-firstIdx+1)

        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            cn := dq.node
            ci := dq.idx
            if cn.Left != nil {
                q = append(q, &qNode{cn.Left, 2*ci})
            }
            if cn.Right != nil {
                q = append(q, &qNode{cn.Right, (2*ci)+1})
            }
            qSize--
        }

    }
    return maxWidth
}