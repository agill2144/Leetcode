
// there can be multiple foods, foods do not depend on anything
// we can go from n foods to 1 person in a BFS manner
// each BFS level == 1 step
func getFood(grid [][]byte) int {
    if grid == nil {return -1}
    q := [][]int{}
    m := len(grid)
    n := len(grid[0])
    destRow := -1
    destCol := -1
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '#' {
                q = append(q, []int{i,j})
            } else if grid[i][j] == '*' {
                destRow = i
                destCol = j
            }
        }
    }
    dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
    level := 0
    // BFS from food to destRow,destCol
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if dq[0] == destRow && dq[1] == destCol {
                return level
            }
            for _, dir := range dirs {
                r := dq[0] + dir[0]
                c := dq[1] + dir[1]
                if r >= 0 && r < m && c >= 0 && c < n && grid[r][c] != 'X' {
                    grid[r][c] = 'X'
                    q = append(q, []int{r,c})
                }
            }
            qSize--
        }
        level++
        
    }
    return -1
}