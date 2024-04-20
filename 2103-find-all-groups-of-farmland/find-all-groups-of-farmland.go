func findFarmland(land [][]int) [][]int {
    m := len(land)
    n := len(land[0])
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    maxR, maxC := math.MinInt64, math.MinInt64
    var dfs func(r, c int)
    dfs = func(r, c int) {
        // base
        if r < 0 || r == m || c < 0 || c == n || land[r][c] != 1 {return}

        // logic
        maxR = max(r, maxR)
        maxC = max(c, maxC)
        land[r][c] = 0
        for _, dir := range dirs {
            dfs(r+dir[0], c+dir[1])
        }
    }
    out := [][]int{}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if land[i][j] == 1 {
                sr, sc := i, j
                maxR, maxC = math.MinInt64, math.MinInt64
                dfs(i,j)
                out = append(out, []int{sr,sc,maxR,maxC})
            }
        }
    }
    return out
}