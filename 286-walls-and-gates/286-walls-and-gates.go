func wallsAndGates(rooms [][]int)  {
    m := len(rooms)
    n := len(rooms[0])
    dirs := [][]int{
        {1,0},
        {-1,0},
        {0,-1},
        {0,1},
    }
    q := [][]int{}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if rooms[i][j] == 0 {
                q = append(q, []int{i,j})
            }
        }
    }
    
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            
            currRow := dq[0]
            currCol := dq[1]
            
            
            for _, dir := range dirs {
                r := currRow + dir[0]
                c := currCol + dir[1]
                if r >= 0 && r < m && c >= 0 && c < n && rooms[r][c] == 2147483647 {
                    rooms[r][c] = rooms[currRow][currCol]+1
                    q = append(q, []int{r,c})
                }
            }
            qSize--
        }
    }
}