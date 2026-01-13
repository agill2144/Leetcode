func spiralOrder(matrix [][]int) []int {
    m := len(matrix)
    n := len(matrix[0])
    t,b := 0, m-1
    l,r := 0, n-1
    out := []int{}
    for len(out) != m*n {
        
        // top row: left to right
        for j := l; j <= r; j++ {
            out = append(out, matrix[t][j])
        }
        t++
        if len(out) == m*n {break}

        // right col: top to bottom
        for i := t; i <= b; i++ {
            out = append(out, matrix[i][r])
        }
        r--
        if len(out) == m*n {break}

        // bottom row: right to left
        for j := r; j >= l; j-- {
            out = append(out, matrix[b][j])
        }
        b--
        if len(out) == m*n {break}

        // left col: bottom to top
        for i := b; i >= t; i-- {
            out = append(out, matrix[i][l])
        }
        l++
    }
    return out
}