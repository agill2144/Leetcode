func islandPerimeter(grid [][]int) int {
    q := [][]int{}
    dirs := [][]int{{1,0},{-1,0},{0,-1}, {0,1}}
    m := len(grid)
    n := len(grid[0])
    out := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                q = append(q, []int{i,j})
                grid[i][j] = -1
                for len(q) != 0 {
                    dq := q[0]
                    q = q[1:]
                    out += 4
                    cr := dq[0]
                    cc := dq[1]
                    for _, dir := range dirs {
                        r := cr + dir[0]
                        c := cc + dir[1]
                        if r >= 0 && r < m && c >= 0 && c < n && (grid[r][c] == 1 || grid[r][c] == -1) {
                            out -= 1
                            if grid[r][c] == 1 {
                                q = append(q, []int{r,c})
                                grid[r][c] = -1
                            }
                        }
                    }
                }
                
            }
        }
    }
    return out
}