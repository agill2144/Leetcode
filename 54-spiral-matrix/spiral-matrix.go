func spiralOrder(matrix [][]int) []int {
    m := len(matrix)
    n := len(matrix[0])
    top := 0
    bottom := m-1
    left := 0
    right := n-1
    out := []int{}
    for len(out) != m*n {
        // top : left -> right (col)
        for j := top; j <= right; j++ {
            out = append(out, matrix[top][j])
        }
        top++
        if len(out) == m*n {break}

        // right : top to bottom ( row )
        for i := top; i <= bottom; i++ {
            out = append(out, matrix[i][right])
        }
        right--
        if len(out) == m*n {break}

        // bottom : right to left ( col )
        for j := right; j >= left; j-- {
            out = append(out, matrix[bottom][j])
        }
        bottom--
        if len(out) == m*n {break}

        // left : bottom to top ( row )
        for i := bottom; i >= top; i-- {
            out = append(out, matrix[i][left])
        }
        left++
    }
    return out
}