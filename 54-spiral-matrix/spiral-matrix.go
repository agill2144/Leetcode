func spiralOrder(matrix [][]int) []int {
    m := len(matrix)
    n := len(matrix[0])
    top := 0
    bottom := m-1
    left := 0
    right := n-1
    out := []int{}
    for len(out) != m*n {

        // left to right , same row
        for i := left; i <= right; i++ {
            out = append(out, matrix[top][i])
        }
        top++
        if len(out) == m*n {break}
        
        // top to bottom on right side
        for i := top; i <= bottom; i++ {
            out = append(out, matrix[i][right])
        }
        right--
        if len(out) == m*n {break}

        // right to left on bottom side
        for i := right; i >= left; i-- {
            out = append(out, matrix[bottom][i])
        }
        bottom--
        if len(out) == m*n {break}

        // bottom to top on left side
        for i := bottom; i >= top; i-- {
            out = append(out, matrix[i][left])
        }
        left++
    }
    return out

}