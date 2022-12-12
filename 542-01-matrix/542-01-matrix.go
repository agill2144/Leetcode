func updateMatrix(mat [][]int) [][]int {
    m := len(mat)
    n := len(mat[0])
    dirs := [][]int{
        {1,0},
        {-1,0},
        {0,1},
        {0,-1},
    }
    q := [][]int{}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if mat[i][j] == 0 {
                q = append(q, []int{i,j})
            } else {
                mat[i][j] = -1
            }
        }
    }
    
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            cr := dq[0]
            cc := dq[1]
            for _, dir := range dirs {
                r := cr + dir[0]
                c := cc + dir[1]
                if r >=0 && r < m && c >= 0 && c < n && mat[r][c] == -1 {
                    mat[r][c] = mat[cr][cc] + 1
                    q = append(q, []int{r,c})
                }
            }
            qSize--
        }
    }
    
    return mat
}