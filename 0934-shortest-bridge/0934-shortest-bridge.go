// classic dfs and bfs
// identify the first island and mark it visited using dfs and at same time add its cords into a queue
// using the queue, perform level order to find the shortest path to nearest 1
func shortestBridge(grid [][]int) int {
    n := len(grid)
    dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
    q := [][]int{}
    var dfs func(r, c int)
    dfs = func(r, c int) {
        // base
        if r < 0 || r == n || c < 0 || c == n || grid[r][c] != 1 {return}
        
        // logic
        q = append(q, []int{r,c})
        grid[r][c] = -1
        for _, dir := range dirs {
            dfs(r+dir[0], c+dir[1])
        }
    }
    for i := 0; i < n; i++ {
        exit := false
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                dfs(i, j)
                exit = true
                break
            }
        }
        if exit {break}
    }
    
    level := 0
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            cr := dq[0]
            cc := dq[1]
            if grid[cr][cc] == 1 {return level-1}
            for _, dir := range dirs {
                nr := dir[0] + cr
                nc := dir[1] + cc
                if nr >= 0 && nr < n && nc >= 0 && nc < n {
                    if grid[nr][nc] == 0 {
                        grid[nr][nc] = -1
                        q = append(q, []int{nr,nc})
                    } else if grid[nr][nc] == 1 {
                        return level
                    }
                }
            }
            qSize--
        }
        level++
    }
    return -1
}