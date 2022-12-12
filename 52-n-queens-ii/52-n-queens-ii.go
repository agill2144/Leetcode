func totalNQueens(n int) int {
    board := make([][]bool, n)
    for idx, _ := range board {
        board[idx] = make([]bool, n)
    }
    count := 0
    var backtrack func(r int)
    backtrack = func(r int) {
        // base
        if r == len(board) {
            count++
            return
        }
        
        // logic
        for i := 0; i < len(board[0]); i++ {
            if isSafe(r,i, board) {
                board[r][i] = true
                backtrack(r+1)
                board[r][i] = false
            }
        }
    }
    backtrack(0)
    return count
}

func isSafe(r,c int, board [][]bool) bool {
    m := len(board)
    n := len(board[0])
    dirs := [][]int{
        {-1,-1},
        {-1,0},
        {-1,1},
    }
    for _, dir := range dirs {
        nr := r + dir[0]
        nc := c + dir[1]
        for nr >= 0 && nr < m && nc >= 0 && nc < n {
            if !board[nr][nc] { 
                nr += dir[0]
                nc += dir[1]   
            } else {
                return false
            }
        }
    }
    return true
}