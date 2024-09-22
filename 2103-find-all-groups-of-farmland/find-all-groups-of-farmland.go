func findFarmland(land [][]int) [][]int {
    m := len(land)
    n := len(land[0])
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    maxR, maxC := -1,-1
    var dfs func(r, c int)
    dfs = func(r, c int) {
        // base
        if r < 0 || r == m || c < 0 || c == n || land[r][c] != 1 {return}

        // logic
        // visit
        land[r][c] = 0
        maxR, maxC = max(maxR, r), max(maxC, c)
        for _, dir := range dirs {
            dfs(r+dir[0], c+dir[1])
        }
    }
    out := [][]int{}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if land[i][j] == 1 {
                dfs(i, j)
                out = append(out, []int{i,j,maxR, maxC})
                maxR, maxC = -1,-1
            }
        }
    }
    return out
}