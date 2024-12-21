func findDiagonalOrder(mat [][]int) []int {
    m := len(mat)
    n := len(mat[0])
    out := []int{}
    r, c := 0, 0
    up := true
    for len(out) != m*n {
        out = append(out, mat[r][c])
        if up {
            // when going up, we could hit
            // top row edge, or right col edge
            if r == 0 || c == n-1 {
                // continue in next col if we are not yet at n-1 col
                // when we are at n-1 col, start from next row, c++ will go out of bounds
                if c != n-1 { c++ } else { r++ }
                up = false
            } else {
                // not hitting edges yet
                // keep going up in this dir
                r--; c++
            }
        } else {
            if c == 0 || r == m-1 {
                if r != m-1 { r++ } else { c++ }
                up = true
            } else {
                // not hitting edges yet
                // keep going up in this dir
                r++; c--
            }
        }
    }
    return out
}