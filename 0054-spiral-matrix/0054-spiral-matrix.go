func spiralOrder(matrix [][]int) []int {
    m := len(matrix)
    n := len(matrix[0])
    out := []int{}
    left := 0
    right := n-1
    top := 0
    bottom := m-1

    
    for len(out) != m*n {
        for i := left; i <= right; i++ {
            out = append(out, matrix[top][i])
        }
        top++
        if len(out) == m*n{break}
        
        for i := top; i <= bottom; i++ {
            out = append(out, matrix[i][right])
        }
        right--
        if len(out) == m*n{break}
        
        for i := right; i >= left; i-- {
            out = append(out, matrix[bottom][i])
        }
        bottom--
        
        if len(out) == m*n{break}
        for i := bottom; i >= top; i-- {
            out = append(out, matrix[i][left])
        }
        left++
    }
    return out
} 