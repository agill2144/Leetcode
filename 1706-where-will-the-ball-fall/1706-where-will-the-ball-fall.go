func findBall(grid [][]int) []int {
    m := len(grid)
    n := len(grid[0])
    // 1  -> \
    // -1 -> /
    
    var follow func(r, c int) (bool, int)
    follow = func(r, c int) (bool, int) {
        for r >= 0 && r < m && c >= 0 && c < n {
            currDir := grid[r][c]
            if currDir == 1 {
                if (c+1 < n && grid[r][c+1] == -1) || c+1 == n {return false, -1}
                r++; c++;
            } else if currDir == -1 {
                if (c-1 >= 0 && grid[r][c-1] == 1) || c-1 < 0 {return false,-1}
                r++; c--
            }
        }
        return true,c
    }

    out := make([]int, n)
    for i := 0; i < n ; i++ {
        if ok, cIdx := follow(0, i); ok {
            out[i] = cIdx
        } else if !ok {
            out[i] = -1
        }
    }
    return out
}

