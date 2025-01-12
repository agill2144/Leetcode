func shortestDistance(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    dists := make([][][]int, m)
    for i := 0; i < m; i++ {
        dists[i] = make([][]int, n)
        for j := 0; j < n; j++ {
            dists[i][j] = make([]int, 2) // totalDist, houseCount
        }
    }
    
    count := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                count++
                bfs(grid, i, j, dists)
            }
        }
    }
    minDist := math.MaxInt64
    for i := 0; i < len(dists); i++ {
        for j := 0; j < len(dists[0]); j++ {
            if dists[i][j][1] == count {
                minDist = min(minDist, dists[i][j][0])
            }
        }
    }
    if minDist == math.MaxInt64 {minDist = -1}
    return minDist
}

func bfs(grid [][]int, r, c int, dists [][][]int) {
    m := len(grid)
    n := len(grid[0])
    visited := make([][]bool,m)
    for i := 0; i < m; i++ {visited[i] = make([]bool, n)}
    q := [][]int{{r,c,0}}
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    visited[r][c] = true
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        cr, cc, cd := dq[0], dq[1], dq[2]
        for _, dir := range dirs {
            nr := cr + dir[0]
            nc := cc + dir[1]
            nd := cd + 1
            if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] == 0 && !visited[nr][nc] {
                visited[nr][nc] = true
                dists[nr][nc][0] += nd
                dists[nr][nc][1]++
                q = append(q, []int{nr,nc,nd})
            }
        }
    }
}