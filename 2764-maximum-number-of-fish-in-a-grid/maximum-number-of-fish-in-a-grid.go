// find out which connected component results into max sum
func findMaxFish(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    landCell := 0    
    dirs := [][]int{
        {0,1},{0,-1},
        {1,0},{-1,0},
    }
    var dfs func(r, c int) int
    dfs = func(r, c int) int  {
        // base
        if r < 0 || r == m || c < 0 || c == n || grid[r][c] == landCell {return 0}
        // logic
        total := grid[r][c]
        grid[r][c] = landCell
        for _ , dir := range dirs {
            total += dfs(r+dir[0], c+dir[1])
        }
        return total
    }
    res := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] != landCell {
                res = max(res, dfs(i, j))
            }
        }
    }
    return res
}