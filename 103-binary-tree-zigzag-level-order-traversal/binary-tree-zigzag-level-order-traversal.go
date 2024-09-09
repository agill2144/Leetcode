/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// level order using bfs with pre-processing of a level
// time = o(2n) = o(n)
// space = o(2n) (queue & levelNodes ) = o(n) 
func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil {return nil}
    out := [][]int{}
    q := []*TreeNode{root}
    level := 1
    for len(q) != 0 {
        // pre-process and get the nodes in the order we care about
        // based on the current level
        // if level is odd; we want left -> right
        // if level is even; we want right -> left
        qSize := len(q)
        dir := 1
        ptr := 0
        if level % 2 == 0 {
            dir = -1
            ptr = qSize-1
        }
        levelNodes := []int{}
        for qSize != 0 {
            levelNodes = append(levelNodes, q[ptr].Val)
            ptr += dir
            qSize--
        }
        out = append(out, levelNodes)
        qSize = len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if dq.Left != nil {q = append(q, dq.Left)}
            if dq.Right != nil {q = append(q, dq.Right)}
            qSize--
        }
        level++
    }
    return out
}