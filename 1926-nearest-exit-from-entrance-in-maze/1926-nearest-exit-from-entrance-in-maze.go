func nearestExit(maze [][]byte, entrance []int) int {
    
    sr := entrance[0]
    sc := entrance[1]
    m := len(maze) // rows
    n := len(maze[0]) // cols
    dirs := [][]int{
        {1,0},
        {-1,0},
        {0,1},
        {0,-1},
    }
    
    q := [][]int{
        {sr,sc},
    }
    maze[sr][sc] = '+'
    level := 0
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            currRow := dq[0]
            currCol := dq[1]
            // fmt.Println("row: ", currRow, " currCol: ", currCol)

            if currRow != sr || currCol != sc{
                if currRow == 0 || currRow == m-1 || currCol == 0 || currCol == n-1 {
                    return level
                }
            }
            
            for _, dir := range dirs {
                r := currRow + dir[0]
                c := currCol + dir[1]
                if r >= 0 && r < m && c >= 0 && c < n && maze[r][c] == '.' {
                    q = append(q, []int{r,c})
                    maze[r][c] = '+'
                }
            }
            qSize--
        }
        level++
    }
    return -1
}