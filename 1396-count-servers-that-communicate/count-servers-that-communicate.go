func countServers(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    rowHasServers := map[int]int{}
    colHasServers := map[int]int{}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            rowHasServers[i] += grid[i][j]
            colHasServers[j] += grid[i][j]
        }
    }
    total := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                if rowHasServers[i] > 1 || colHasServers[j] > 1 {
                    total++
                }
            }
        }
    }
    return total
}