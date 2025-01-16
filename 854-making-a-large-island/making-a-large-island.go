func largestIsland(grid [][]int) int {
    dirs := [][]int{
        {-1,0},
        {1,0},
        {0,1},
        {0,-1},
    }
    n := len(grid)
    idToSize := map[int]int{}
    var dfs func(r, c, id int) int
    dfs = func(r, c, id int) int {
        // base
        if r < 0 || r == n || c < 0 || c == n || grid[r][c] != 1 {return 0}

        // logic
        size := 1
        grid[r][c] = id
        for _, dir := range dirs {
            size += dfs(r+dir[0],c+dir[1], id)
        }
        return size
    }
    id := 2
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                idToSize[id] = dfs(i,j,id)
                id++
            }
        }
    }
    res := 0
    haveZero := false
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 0 {
                haveZero = true
                uniqIDs := map[int]bool{}
                for _, dir := range dirs {
                    nr := i+dir[0]
                    nc := j+dir[1]
                    if nr >= 0 && nr < n && nc >= 0 && nc < n && grid[nr][nc] != 0 {
                        uniqIDs[grid[nr][nc]] = true
                    }
                }
                totalSize := 1
                for id, _ := range uniqIDs {
                    totalSize += idToSize[id]
                }
                res = max(res, totalSize)
            }
        }
    }
    if !haveZero {res = n*n}
    return res
}