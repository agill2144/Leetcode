func countNegatives(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    r := m-1
    c := 0
    total := 0
    for r >= 0 && c < n {
        if grid[r][c] < 0 {
            total += (n-c)
            r--
        } else {
            c++
        }
    }
    return total
}