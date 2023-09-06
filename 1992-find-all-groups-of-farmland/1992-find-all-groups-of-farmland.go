/*
    find the cordinates for each connected component
    - we need to know where a connected component starts and ends at ( row and col index for each )
*/
func findFarmland(land [][]int) [][]int {
    m := len(land)
    n := len(land[0])
    out := [][]int{}
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    
    er, ec := -1,-1
    var dfs func(r, c int)
    dfs = func(r, c int) {
        // base
        if r < 0 || r == m || c < 0 || c == n || land[r][c] != 1 {return}
        
        
        // logic
        if r > er {er=r}
        if c > ec {ec=c}
        land[r][c] = 0
        for _, dir := range dirs {
            dfs(r+dir[0], c+dir[1])
        }
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if land[i][j] == 1 {
                dfs(i,j)
                out = append(out, []int{i,j,er,ec})
                er,ec = -1,-1
            }
        }
    }
    return out
}