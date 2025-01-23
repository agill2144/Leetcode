// got confused , looked like a connected compoenent problem at first
// but its not
// a server to be counted as connected, its row or col must have atleast 1 more server excluding itself
// hence we track server freq by row and col in a map
// then take another pass to count total servers
// tc = o(mn) + o(mn)
// sc = o(m+n)
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