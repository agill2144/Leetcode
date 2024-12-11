func largestIsland(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    maxArea := 0
    idToArea := map[int]int{}
    var dfs func(r, c, id int)
    dfs = func(r, c, id int) {
        // base
        if r == m || r < 0 || c < 0 || c == n || grid[r][c] != 1 {return}

        // logic
        grid[r][c] = id
        idToArea[id]++
        maxArea = max(maxArea, idToArea[id])
        for _, dir := range dirs {
            dfs(r+dir[0], c+dir[1], id)
        }
    }
    id := 2
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                dfs(i, j, id)
                id++
            }
        }
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 0 {
                uniqIds := map[int]bool{}
                for _, dir := range dirs {
                    nr := i+dir[0]
                    nc := j+dir[1]
                    if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] != 0 {
                        uniqIds[grid[nr][nc]] = true
                    }
                }

                area := 1
                for i, _ := range uniqIds {
                    area += idToArea[i] 
                }
                maxArea = max(maxArea, area)
            }
        }
    }
    return maxArea
}