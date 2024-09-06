func countNegatives(grid [][]int) int {
    count := 0
    m := len(grid)
    n := len(grid[0])
    r := m-1
    c := 0
    for r >= 0 && c < n {
        if grid[r][c] < 0 {
            count += (n-c)
            r--
        } else {
            c++
        }
    }
    return count
}