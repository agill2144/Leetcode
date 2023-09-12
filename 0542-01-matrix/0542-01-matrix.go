func updateMatrix(mat [][]int) [][]int {
    m := len(mat)
    n := len(mat[0])
    q := [][]int{}
    for i := 0; i < m ;i++ {
        for j := 0; j < n; j++ {
            if mat[i][j] == 0 {
                q = append(q, []int{i,j})
            } else {
                mat[i][j] = -1
            }
        }
    }
    if len(q) == 0 {return mat}
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}    
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        cr := dq[0]
        cc := dq[1]
        for _, dir := range dirs {
            nr := cr+dir[0]
            nc := cc+dir[1]
            if nr >= 0 && nr < m && nc >= 0 && nc < n && mat[nr][nc] == -1 {
                mat[nr][nc] = mat[cr][cc] + 1
                q = append(q, []int{nr, nc})
            }  
        }
    }   
    return mat
}