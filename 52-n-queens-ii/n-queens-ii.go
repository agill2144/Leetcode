func totalNQueens(n int) int {
    count := 0
    board := make([][]byte, n)
    for i := 0; i < n; i++ {board[i] = make([]byte, n)}
    var dfs func(row int)
    dfs = func(row int) {
        // base
        if row == n {
            count++
            return
        }

        // logic
        for j := 0; j < n; j++ {
            if canPlace(board, row, j) {
                board[row][j] = 'Q'
                dfs(row+1)
                board[row][j] = '.'
            }
        }
    }
    dfs(0)
    return count
}


func canPlace(board [][]byte, row, col int) bool {
    n := len(board)
    dirs := [][]int{{-1,0},{-1,-1},{-1,1}}
    for _, dir := range dirs {
        nr := row+dir[0]
        nc := col+dir[1]
        for nr >= 0 && nc >= 0 && nc < n {
            if board[nr][nc] == 'Q' {return false}
            nr += dir[0]
            nc += dir[1]
        }
    }
    return true
}