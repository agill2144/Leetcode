func updateMatrix(matrix [][]int) [][]int {
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    m := len(matrix)
    n := len(matrix[0])

    visited := make([][]bool, m)
    for i := 0; i < m; i++ {visited[i] = make([]bool, n)}
    
    q := [][]int{}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0 {
                q = append(q, []int{i,j})
                visited[i][j] = true
            }
        }
    }
    
    
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        cr := dq[0]
        cc := dq[1]
        currDist := matrix[cr][cc]
        for _, dir := range dirs {
            nr := cr + dir[0]
            nc := cc + dir[1]
            if nr >= 0 && nr < m && nc >= 0 && nc < n && !visited[nr][nc]{
                matrix[nr][nc] = currDist+1
                q = append(q, []int{nr,nc})
                visited[nr][nc] = true
            }
        }
    }
    return matrix
}