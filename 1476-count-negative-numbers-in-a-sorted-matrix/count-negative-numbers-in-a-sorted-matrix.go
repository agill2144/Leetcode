func countNegatives(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    i := m-1
    j := 0
    count := 0
    for i >= 0 && j < n {
        if grid[i][j] < 0 {
            count += (n-j)
            i--
        } else {
            j++
        }
    }
    return count
}