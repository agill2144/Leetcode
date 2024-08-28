// we want to know whether islands in grid2 are also islands in grid1
// use dfs to find connected lands in grid2
// and keep making sure that grid1 at specific r,c is also a land
// if a pariticular land does not exist in grid1, then turn some flag that indicates
// this specific connected-land does not exist in grid1, therefore not a subisland
// therefore only count when that state/flag says so
// time = o(mn)
// space = o(1)
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
    m := len(grid1)
    n := len(grid1[0])    
    dirs := [][]int{
        {1,0},{-1,0},
        {0,1},{0,-1},
    }
    
    isSubIsland := true
    var dfs func(r, c int) 
    dfs = func(r, c int) {
        // base
        if r < 0 || r == m || c < 0 || c == n || grid2[r][c] != 1 {return}

        // logic
        if grid1[r][c] != 1 {
            isSubIsland = false
        }
        grid1[r][c] = 0
        grid2[r][c] = 0
        for _, dir := range dirs {
            dfs(r+dir[0], c+dir[1])
        }
    }
    count := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid2[i][j] == 1 {
                dfs(i, j)
                if isSubIsland {count++}
                isSubIsland = true
            }
        }
    }
    return count
}