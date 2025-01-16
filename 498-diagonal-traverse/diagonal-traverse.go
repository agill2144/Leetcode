func findDiagonalOrder(mat [][]int) []int {
    m := len(mat)
    n := len(mat[0])
    r := 0
    c := 0
    out := []int{}
    up := true
    for len(out) != m*n {
        out = append(out, mat[r][c])
        if up {
            if c == n-1 {
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
            if r == m-1 {
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