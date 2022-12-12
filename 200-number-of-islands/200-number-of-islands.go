func numIslands(grid [][]byte) int {
    
    m := len(grid)
    n := len(grid[0])
    dirs := [][]int{
        {1,0},
        {-1,0},
        {0,-1},
        {0,1},
    }
    islands := 0
    q := [][]int{}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' {
                grid[i][j] = '2'
                q = append(q, []int{i,j})
                
                for len(q) != 0 {
                    dq := q[0]
                    q = q[1:]
                    cr := dq[0]
                    cc := dq[1]
                    
                    for _, dir := range dirs {
                        r := cr + dir[0]
                        c := cc + dir[1]
                        if r >= 0 && r < m && c >= 0 && c < n && grid[r][c] == '1' {
                            grid[r][c] = '2'
                            q = append(q, []int{r,c})
                        }
                    }
                }
                islands++
            }
        }
    }
    
    return islands
    
    
}