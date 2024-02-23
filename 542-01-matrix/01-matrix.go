func updateMatrix(mat [][]int) [][]int {
    m := len(mat)
    n := len(mat[0])
    q := [][]int{}
    visited := make([][]bool, m)
    for i := 0; i < m ; i++ {
        visited[i] = make([]bool, n)
        for j := 0; j < n; j++ {
            if mat[i][j] == 0 {
                q = append(q, []int{i,j})
                visited[i][j] = true
            }
        }
    }
    if len(q) == 0 {return mat}
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    level := 0
    for len(q) != 0 {
        qSize := len(q) 
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            cr := dq[0]
            cc := dq[1]
            for _, dir := range dirs {
                nr := cr + dir[0]
                nc := cc + dir[1]
                if nr >= 0 && nr < m && nc >= 0 && nc < n && mat[nr][nc] == 1 && !visited[nr][nc] {
                    mat[nr][nc] = level+1
                    visited[nr][nc] = true
                    q = append(q, []int{nr,nc})
                }
            }
            qSize--
        }
        level++
    }    
    return mat
}