func largestIsland(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    idToSize := map[int]int{}
    dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
    var dfs func(r, c, id int) int
    dfs = func(r, c, id int) int {
        // base
        if r < 0 || r == m || c < 0 || c == n || grid[r][c] != 1 {return 0}

        // logic
        total := 1
        grid[r][c] = id
        for _, dir := range dirs {
            total += dfs(r+dir[0], c+dir[1], id)
        } 
        return total
    }

    id := 2
    haveZeros := false
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                idToSize[id] = dfs(i, j, id)
                id++
            } else if grid[i][j] == 0 {
                haveZeros = true
            }
        }
    }
    if !haveZeros { return m*n }

    res := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 0 {
                uniqIds := map[int]bool{}
                for _, dir := range dirs {
                    nr, nc := i+dir[0], j+dir[1]
                    if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] != 0 {
                        uniqIds[grid[nr][nc]] = true
                    }
                }
                size := 1
                for k, _ := range uniqIds {
                    size += idToSize[k]
                }
                res = max(res, size)
            }
        }
    }
    return res
}
