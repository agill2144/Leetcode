/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {
    q := [][]interface{}{{root, 0, ""}}
    max := 0
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        node := dq[0].(*TreeNode)
        size := dq[1].(int)
        dir := dq[2].(string)
        if size > max {max = size}
        
        if node.Left != nil {
            newLeftSize := 1
            if dir == "" || dir == "R" {
                newLeftSize += size
            }
            q = append(q, []interface{}{node.Left, newLeftSize, "L"})
        }
        if node.Right != nil {
            newRightSize := 1
            if dir == "" || dir == "L" {
                newRightSize += size
            }
            q = append(q, []interface{}{node.Right, newRightSize, "R"})            
        }
        
    }
    
    
    return max
}