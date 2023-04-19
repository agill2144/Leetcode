/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    keep track of which side per node we are coming from
    and keep track of number of steps at each node.
    
    when current dq'd node came from left
    and this node has a left child, this is a straight path
    therefore enqueue left child with steps resetted to 1
    which also means right child of current node is a continuation of a zigzag,
    therefore enqueue right child with current steps + 1
    
    time; o(n)
    space: o(n)
    
*/
func longestZigZag(root *TreeNode) int {
    q := [][]interface{}{{root, 0, 0}}
    // 1 = right
    // -1 = left
    max := 0
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        node := dq[0].(*TreeNode)
        size := dq[1].(int)
        dir := dq[2].(int)
        if size > max {max = size}
        
        if node.Left != nil {
            newLeftSize := 1
            if dir == 0 || dir == 1 {
                newLeftSize += size
            }
            q = append(q, []interface{}{node.Left, newLeftSize, -1})
        }
        if node.Right != nil {
            newRightSize := 1
            if dir == 0 || dir == -1 {
                newRightSize += size
            }
            q = append(q, []interface{}{node.Right, newRightSize, 1})
        }
    }
    
    
    return max
}