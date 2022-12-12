/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func spiralMatrix(m int, n int, head *ListNode) [][]int {
    if head == nil {return nil}
    matrix := make([][]int, m)
    for i := 0; i < m; i++ {
        matrix[i] = make([]int, n)
        for j := 0; j < n; j++{
            matrix[i][j] = -1
        }
    }
    
    current := head
    left := 0
    right := n-1
    top := 0
    bottom := m-1
    for current != nil {
        for i := left; i <= right && current != nil; i++ {
            matrix[top][i] = current.Val
            current = current.Next
        }
        top++
        
        for i := top; i <= bottom && current != nil; i++ {
            matrix[i][right] = current.Val
            current = current.Next
        }
        right--
        
        for i := right; i >= left && current != nil; i-- {
            matrix[bottom][i] = current.Val
            current = current.Next
        }
        bottom--
        
        for i := bottom; i >= top && current != nil; i-- {
            matrix[i][left] = current.Val
            current = current.Next
        }
        left++
    }
    return matrix
}
