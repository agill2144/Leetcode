/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func replaceValueInTree(root *TreeNode) *TreeNode {
    if root == nil {return root}
    
    type qNode struct {
        node *TreeNode
        siblingSum int
    }
    q := []*qNode{&qNode{root,root.Val}}
    levelSum := root.Val
    for len(q) != 0 {
        nextLevelSum := 0
        qSize := len(q) 
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            node := dq.node
            node.Val = levelSum-dq.siblingSum
            
            nextSiblingSum := 0
            if node.Left != nil {nextSiblingSum += node.Left.Val}
            if node.Right != nil {nextSiblingSum += node.Right.Val}
            nextLevelSum += nextSiblingSum
            
            if node.Left != nil {
                q = append(q, &qNode{node.Left, nextSiblingSum})
            }
            if node.Right != nil {
                q = append(q, &qNode{node.Right, nextSiblingSum})
            }
            qSize--
        }
        levelSum = nextLevelSum
    }
    return root
}