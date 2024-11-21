func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
    if len(guards) == 0 { return (m*n)-len(walls)  }
    wall := 1
    guard := 2
    visited := 3
    matrix := make([][]int, m)
    total := m*n
    for i := 0; i < m; i++ {matrix[i] = make([]int, n)}
    for i := 0; i < len(walls); i++ {
        r, c := walls[i][0], walls[i][1]
        matrix[r][c] = wall
        total--
    }
    q := [][]int{}
    for i := 0; i < len(guards); i++ {
        r, c := guards[i][0], guards[i][1]
        matrix[r][c] = guard
        q = append(q, guards[i])
        total--
    }
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        cr := dq[0]
        cc := dq[1]
        for _, dir := range dirs {
            nr := cr + dir[0]
            nc := cc + dir[1]
            for nr >= 0 && nr < m && nc >= 0 && nc < n && (matrix[nr][nc] != guard && matrix[nr][nc] != wall) {
                if matrix[nr][nc] != visited {total--}
                matrix[nr][nc] = visited
                nr += dir[0]
                nc += dir[1]                
            }       
        }
    }
    
    /*
        wall := 1
        guard := 2
        visited := 3

    
        [2 1 0 0 0 0] 
        [3 2 3 3 1 0] 
        [3 3 1 2 3 3] 
        [3 3 0 3 0 0]
    */
    return total
}