func solveSudoku(board [][]byte) {
    var rows, cols [9][10]bool
    var boxes [3][3][10]bool

    // First pass: mark existing numbers
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] != '.' {
                num := int(board[i][j] - '0')
                rows[i][num] = true
                cols[j][num] = true
                boxes[i/3][j/3][num] = true
            }
        }
    }

    var solve func(int, int) bool
    solve = func(r, c int) bool {
        // Base case: solved the entire board
        if r == 9 {
            return true
        }

        // Move to next row if column is full
        if c == 9 {
            return solve(r+1, 0)
        }

        // Skip filled cells
        if board[r][c] != '.' {
            return solve(r, c+1)
        }

        // Try placing numbers 1-9
        for num := 1; num <= 9; num++ {
            // Check if number can be placed
            if !rows[r][num] && !cols[c][num] && !boxes[r/3][c/3][num] {
                // Place the number
                board[r][c] = byte(num) + '0'
                rows[r][num] = true
                cols[c][num] = true
                boxes[r/3][c/3][num] = true

                // Try solving the rest of the board
                if solve(r, c+1) {
                    return true
                }

                // Backtrack
                board[r][c] = '.'
                rows[r][num] = false
                cols[c][num] = false
                boxes[r/3][c/3][num] = false
            }
        }

        return false
    }

    solve(0, 0)
}
