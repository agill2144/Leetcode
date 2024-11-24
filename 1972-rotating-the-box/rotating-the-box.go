func rotateTheBox(matrix [][]byte) [][]byte {
    m := len(matrix)
    n := len(matrix[0])
    out := make([][]byte, n)
    for i := 0; i < n; i++ {out[i] = make([]byte, m)}

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            out[j][i] = matrix[i][j]
        }
    }

    for i := 0; i < n; i++ {
        left := 0
        right := m-1
        for left < right {
            out[i][left], out[i][right] = out[i][right], out[i][left]
            left++
            right--
        }
    }

    c := 0
    for c < len(out[0]) {
        ne := len(out)-1
        for ne >= 0 && out[ne][c] != '.' {ne--}
        if ne <= 0 {c++; continue}
        r := ne-1        
        for r >= 0 && ne > 0 {
            if out[r][c] == '#' {
                out[ne][c], out[r][c] = out[r][c], out[ne][c]
                ne--
                for ne >= 0 && out[ne][c] != '.' {ne--}
                r = ne-1
            } else if out[r][c] == '*' {
                ne = r-1
                r--
            } else {
                r--
            }
        }
        c++
    }

    
    return out
}

/*

        ["#","#"]
        ["#","."]
        [".","*"]
r ne    ["*","."]
          c

*/