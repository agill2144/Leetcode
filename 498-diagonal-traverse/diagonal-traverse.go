func findDiagonalOrder(mat [][]int) []int {
    m := len(mat)
    n := len(mat[0])
    out := []int{}
    r, c := 0, 0
    up := true
    for len(out) != m*n {
        out = append(out, mat[r][c])
        if up {
            if r == 0 || c == n-1 {
                if c != n-1 {
                    c++
                } else {
                    r++
                }
                up = false
            } else {
                r--
                c++
            }
        } else {
            if c == 0 || r == m-1 {
                if r != m-1 {
                    r++
                } else {
                    c++
                }
                up = true
            } else {
                r++
                c--
            }
        }
    }
    return out
}