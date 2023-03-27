func minPathSum(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            if i==m-1 && j==n-1 {continue}
            rightCell := math.MaxInt64
            if j != n-1 {rightCell = grid[i][j+1]}
            bottomCell := math.MaxInt64
            if i != m-1 {bottomCell = grid[i+1][j]}
            grid[i][j] += min(rightCell, bottomCell)
        }
    }
    return grid[0][0]
}

func min(x, y int) int {
    if x < y {return x}
    return y
}