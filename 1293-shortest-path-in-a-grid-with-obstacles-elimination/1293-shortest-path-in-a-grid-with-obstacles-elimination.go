func shortestPath(grid [][]int, k int) int {
    
    m := len(grid)
    n := len(grid[0])
    
    visited := make([][][]bool, m)
    for idx, _ := range visited {
        visited[idx] = make([][]bool, n)
        for idxN, _ := range visited[idx] {
            visited[idx][idxN] = make([]bool, k+1)
        }
    }
    
    destRow := m-1
    destCol := n-1
    
    dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
    q := [][]int{{0,0, k}} // r,c,k
    visited[0][0][k] = true
    
    level := 0
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            
            cr := dq[0]
            cc := dq[1]
            ck := dq[2]
            
            if cr == destRow && cc == destCol {
                return level
            }
            
            for _, dir := range dirs {
                r := cr + dir[0]
                c := cc + dir[1]
                if r >= 0 && r < m && c >= 0 && c < n {
                    nk := ck - grid[r][c] // eliminate obstacle if present at this cell
                    if nk >= 0 && !visited[r][c][nk] {
                        q = append(q, []int{r,c,nk})
                        visited[r][c][nk] = true
                    }
                }
            }
            qSize--
        }
        level++
    }
    return -1
    
}