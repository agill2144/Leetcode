func shortestPathBinaryMatrix(grid [][]int) int {
    
    m := len(grid)
    n := len(grid[0])
    destRow := m-1
    destCol := n-1
    
    q := [][]int{
        {0,0},
    }
    if grid[0][0] == 1 {return -1}
    
    dirs := [][]int{
        {0,1},
        {0,-1},
        {1,0},
        {-1,0},
        {-1,1},
        {-1,-1},
        {1,-1},
        {1,1},
    }
    
    // visited
    grid[0][0] = 1
    level := 0
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0] // [0,0]
            q = q[1:]
            currRow := dq[0]
            currCol := dq[1]
            
            if currRow == destRow && currCol == destCol {
                return level+1
            }
            
            for _, dir := range dirs {
                r := currRow + dir[0]
                c := currCol + dir[1]
                if r >= 0 && r < m && c >= 0 && c < n && grid[r][c] == 0 {
                    grid[r][c] = 1 // mark it visited
                    q = append(q, []int{r,c})
                }
            }
            qSize--
        }
        level++
    }
    return -1
}