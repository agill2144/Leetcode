func islandPerimeter(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    q := [][]int{}
    total := 0
    visited := 2
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                q = append(q, []int{i,j})
                grid[i][j] = visited
                for len(q) != 0 {
                    dq := q[0]
                    q = q[1:]
                    cr := dq[0]
                    cc := dq[1]
                    count := 4 // this cell has 4 boundaries
                    for _, dir := range dirs {
                        nr := cr + dir[0]
                        nc := cc + dir[1]
                        if nr >= 0 && nr < m && nc >= 0 && nc < n {
                            if grid[nr][nc] == 1 || grid[nr][nc] == visited {
                                // current has a adj land
                                // therefore curr land's boundary is decremented by 1
                                count--
                                if grid[nr][nc] == 1 {
                                    // enqueue, and mark it visited
                                    q = append(q, []int{nr, nc})
                                    grid[nr][nc] = visited
                                }
                            }
                        }
                    }
                    total += count
                }
            }
        }
    }
    return total
}