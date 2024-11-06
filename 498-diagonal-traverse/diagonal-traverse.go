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
            if c == n-1 { up = false; r++; continue}
            if r == 0 { up = false; c++; continue}
            r--; c++
        } else {
            if r == m-1 {c++; up = true; continue}
            if c == 0 {r++; up = true; continue}
            r++; c--
        }
    }
    return out
}