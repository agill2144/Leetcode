func numEnclaves(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    
    
    q := [][]int{}
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == 0 || j == 0 || i == m-1 || j == n-1 {
                if grid[i][j] == 1 {
                    q = append(q, []int{i,j})
                    grid[i][j] = -1
                }
            }
        }
    }

    
    
    for len(q) != 0 {
        dq := q[0]; q = q[1:]
        for _, dir := range dirs {
            nr := dq[0] + dir[0]
            nc := dq[1] + dir[1]
            if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] == 1 {
                grid[nr][nc] = -1
                q = append(q, []int{nr,nc})
            }
        }
    }
    count := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                count++
            }
        }
    }

    return count
}

// approach: start dfs from boundary cells and mark all connected components
// time: o(mn)
// space: o(mn)
// func numEnclaves(grid [][]int) int {
//     m := len(grid)
//     n := len(grid[0])
    
//     dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
//     var dfs func(r, c int)
//     dfs = func(r, c int) {
//         // base
//         if r == m || r < 0 || c == n || c < 0 || grid[r][c] != 1 {return}
        
//         // logic
//         grid[r][c] = -1 // marked
//         for _, dir := range dirs {
//             dfs(r+dir[0], c+dir[1])
//         }
//     }
    
//     for i := 0; i < m; i++ {
//         if grid[i][0] == 1 {dfs(i,0)}
//         if grid[i][n-1] == 1 {dfs(i,n-1)}
//     }
    
//     for j := 0; j < n; j++ {
//         if grid[0][j] == 1 {dfs(0,j)}
//         if grid[m-1][j] == 1 {dfs(m-1,j)}
//     }
//     count := 0
//     for i := 0; i < m; i++ {
//         for j := 0; j < n; j++ {
//             if grid[i][j] == 1{count++}
//         }
//     }
//     return count
// }