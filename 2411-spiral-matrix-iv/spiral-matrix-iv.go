/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func spiralMatrix(m int, n int, head *ListNode) [][]int {
    matrix := make([][]int, m)
    for i := 0; i < len(matrix); i++ {
        matrix[i] = make([]int, n)
        for j := 0; j < n; j++ {
            matrix[i][j] = -1
        } 
    }
    top := 0
    bottom := m-1
    left := 0
    right := n-1
    curr := head
    for curr != nil && top <= bottom && left <= right {
        for j := top; j <= right && curr != nil; j++ {
            matrix[top][j] = curr.Val
            curr = curr.Next
        }
        top++
        if curr == nil || top > bottom || left > right {break}

        // right : top to bottom ( row )
        for i := top; i <= bottom && curr != nil; i++ {
            matrix[i][right] = curr.Val
            curr = curr.Next
        }
        right--
        if curr == nil || top > bottom || left > right {break}

        // bottom : right to left ( col )
        for j := right; j >= left && curr != nil; j-- {
            matrix[bottom][j] = curr.Val
            curr = curr.Next
        }
        bottom--
        if curr == nil || top > bottom || left > right {break}

        // left : bottom to top ( row )
        for i := bottom; i >= top && curr != nil; i-- {
            matrix[i][left] = curr.Val
            curr = curr.Next
        }
        left++
    }
    return matrix
}