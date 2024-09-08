/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    given level starts from 1 ( instead of 0 )
    even level needs to have their nodes inserted from right -> left
    using bfs, before processing the queue
    we can check what level we are in and add the current level nodes in the correct order
    then do a traditional bfs processing

*/
func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    out := [][]int{}
    level := 1
    q := []*TreeNode{root}
    for len(q) != 0 {

        // pre-processing of a level 
        // add nodes first!
        qSize := len(q)
        ptr := 0
        addDir := 1
        if level % 2 == 0 {ptr = qSize-1; addDir = -1}
        levelNodes := []int{}
        for qSize != 0 {
            levelNodes = append(levelNodes, q[ptr].Val)
            ptr += addDir
            qSize--
        }
        out = append(out, levelNodes)

        // now classic bfs processing
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
