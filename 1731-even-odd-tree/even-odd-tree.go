/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isEvenOddTree(root *TreeNode) bool {
    // even level = asc order
    // odd level = desc order
    if root == nil {return true}
    q := []*TreeNode{root}
    level := 0
    for len(q) != 0 {
        qSize := len(q)
        prev := math.MinInt64
        for qSize != 0 {
            dq := q[0]
            q = q[1:]

            if level % 2 == 0 {
                // even level = asc order
                // even level = odd integer
                isOdd := dq.Val % 2 != 0
                if !isOdd || (prev != math.MinInt64 && prev >= dq.Val) {
                    return false
                }
                prev = dq.Val
            } else {
                // odd level = desc order
                // odd level = even integer
                isEven := dq.Val % 2 == 0
                if !isEven || (prev != math.MinInt64 && prev <= dq.Val) {
                    return false
                    
                }
                prev = dq.Val
            }
            if dq.Left != nil {q = append(q, dq.Left)}
            if dq.Right != nil {q = append(q, dq.Right)}
            qSize--
        }
        level++
    }
    return true
}