func minimumArea(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    firstRow := m+1
    lastRow := -1
    firstCol := m+1
    lastCol := -1
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                firstRow = min(firstRow, i)
                firstCol = min(firstCol, j)

                lastRow = max(i, lastRow)
                lastCol = max(j, lastCol)
            }
        }
    }
    height := lastRow-firstRow+1
    width := lastCol-firstCol+1
    return height*width
}