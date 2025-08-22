func minimumArea(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    firstRow := m
    firstCol := n
    lastRow := -1
    lastCol := -1
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                firstRow = min(firstRow, i)
                firstCol = min(firstCol, j)
                lastRow = max(lastRow, i)
                lastCol = max(lastCol, j)
            }
        }
    }
    h := lastRow-firstRow+1
    w := lastCol-firstCol+1
    return h*w
}