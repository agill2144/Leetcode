/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func reverseOddLevels(root *TreeNode) *TreeNode {
    if root == nil {return root}
    q := []*TreeNode{root}
    level := 0
    for len(q) != 0 {
        qSize := len(q)
        levelNodes := []*TreeNode{}
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if level % 2 != 0 {levelNodes = append(levelNodes, dq)}
            if dq.Left != nil {q = append(q, dq.Left)}
            if dq.Right != nil {q = append(q, dq.Right)}
            qSize--
        }
        for i := 0; i < len(levelNodes)/2; i++ {
            levelNodes[i].Val, levelNodes[len(levelNodes)-1-i].Val = levelNodes[len(levelNodes)-1-i].Val, levelNodes[i].Val
        }
        level++
    }
    return root
}