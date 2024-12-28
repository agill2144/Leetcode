func findDiagonalOrder(matrix [][]int) []int {
    m := len(matrix)
    n := len(matrix[0])
    r := 0
    c := 0
    up := true
    out := []int{}
    for len(out) < m*n {
        out = append(out, matrix[r][c])
        if up {
            if  c == n-1 {
                r++
                up = false
            } else if r == 0 {
                c++
                up = false
            } else {
                r--
                c++
            }
        } else {
            if r == m-1  {
                c++
                up = true
            } else if c == 0 {
                r++
                up = true
            } else {
                r++
                c--
            }
        }
    }
    return out
}