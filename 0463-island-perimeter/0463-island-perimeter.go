// func islandPerimeter(grid [][]int) int {
//     q := [][]int{}
//     dirs := [][]int{{1,0},{-1,0},{0,-1}, {0,1}}
//     m := len(grid)
//     n := len(grid[0])
//     out := 0
//     for i := 0; i < m; i++ {
//         for j := 0; j < n; j++ {
//             if grid[i][j] == 1 {
//                 q = append(q, []int{i,j})
//                 grid[i][j] = -1
//                 for len(q) != 0 {
//                     dq := q[0]
//                     q = q[1:]
//                     out += 4
//                     cr := dq[0]
//                     cc := dq[1]
//                     for _, dir := range dirs {
//                         r := cr + dir[0]
//                         c := cc + dir[1]
//                         if r >= 0 && r < m && c >= 0 && c < n && (grid[r][c] == 1 || grid[r][c] == -1) {
//                             out -= 1
//                             if grid[r][c] == 1 {
//                                 q = append(q, []int{r,c})
//                                 grid[r][c] = -1
//                             }
//                         }
//                     }
//                 }
                
//             }
//         }
//     }
//     return out
// }


func islandPerimeter(grid [][]int) int {
    dirs := [][]int{{1,0},{-1,0},{0,-1}, {0,1}}
    m := len(grid)
    n := len(grid[0])
    out := 0
    var dfs func(r, c int)
    dfs = func(r, c int) {
        // base
        
        //logic
        grid[r][c] = -1
        out += 4
        for _, dir := range dirs {
            nr := r+dir[0]
            nc := c+dir[1]
            if nr >= 0 && nr < m && nc >= 0 && nc < n && (grid[nr][nc] == -1||grid[nr][nc] == 1) {
                out -= 1
                if grid[nr][nc] == 1 { 
                    dfs(nr, nc)
                }
            }
        }
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1{
                dfs(i,j)
            }
        }
    }
    return out
}