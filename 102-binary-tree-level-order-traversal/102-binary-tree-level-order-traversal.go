/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil {return nil}
    
    q := []*TreeNode{root}
    result := [][]int{}    
    for len(q) != 0 {
        arr := []int{}
        qSize := len(q)
        
        for qSize != 0 {
            dq := q[0]
            q = q[1:]            
            arr = append(arr, dq.Val)
            if dq.Left != nil { q = append(q, dq.Left) }
            if dq.Right != nil { q = append(q, dq.Right) }    
            qSize--
        }
        
        result = append(result, arr)
    }
    return result
}