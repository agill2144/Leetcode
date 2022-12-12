/*
    First we have to create nxn board.
    Then we have to find every single possible combinations that can be used
    to successfully place all n queens on the n board.
    A queen can be successfully placed at a specific location if its safe.
    What determines a location "safe" for a queen to be set into?
    - When that place cannot be hit by any other queen currently placed on the board
    - A queen can move across all cells vertically, horizontally and diagonally too.
    - So as long as the current position in question cannot be reached by existing queens from the board, it is considered safe.
    
    approach:
    - We will first create the nxn board of type bools.
        - Why? inorder to mark a location "taken" by a queen, we will change that default false to true
        - We will then use this information to later eval whether a location is safe or not.
    - Now we need to start placements of queen in each row.
    - Therefore we will run a for loop on all rows
    - If at a specific col in this row we are iterating over is safe, we will change the value of this cell to true
    - Then recursively call the function again with row+1, because we now need to check where in the next row we need to place the next queen
    - Then that recursive function will recursively call itself once it find a specific col safe and so on.
    - If in a row, we could not find a col safe for a queen, we will return false and our recursion will go back to parent row
    - Which means, as soon as a parent is passed back the control from its child row, it needs to undo everything that was done to any
        reference based data structure ( i.e backtrack the placement of the queen )
    - Then that parent for loop would continue, i++, and then find next safe place and call itself with row+1 and then child does the same thing, until we are done placing all the queens.
    - We would know we have successfully placed ALL N queen by 2 ways
        - We could either keep track of remaining queens, and if at some point we have 0 queens remaining, save our current board
        - Or no need to save another parameter at each recursive stack, simply when row parameter goes out of bouds, that itself will be our indicator that we were able to go thru all rows and were able to find a safe spot in each row and therefore now row is out of bounds. Which means, we have placed all of our queens and have a board combination to save.
        
        
    Time: o(n!)
    space: n^2
    
*/

func solveNQueens(n int) [][]string {
    
    // generate the nxn board
    board := make([][]bool, n)
    for idx, _ := range board {
        board[idx] = make([]bool, n)
    }
    
    var result [][]string
    // now try placing queens 
    var backtrack func(row int)
    backtrack = func(row int) {
        // base
        if row == n {
            tmp := []string{}
            for i := 0; i < n; i++ {
                tmpStr := ""
                for j := 0; j < n; j++ {
                    if board[i][j] == true {
                        tmpStr += "Q"
                    } else {
                        tmpStr += "."
                    }
                }
                tmp = append(tmp, tmpStr)
            }
            result = append(result, tmp)
            return
        }
        
        //logic
        // we loop over cols, row is controlled by recursion
        for i := 0; i < n; i++ {
            if canBePlaced(row, i, board) {
                // action
                board[row][i] = true
                // recurse
                backtrack(row+1)
                // backtrack
                board[row][i] = false
            }
        }
        
    }
    backtrack(0)
    return result
}


// since we are looping row by row
// the only directions we have to check for a specific [r][c]
// are vertical up, diagonal up left and diagonal up right
// we do not have to check in the same row because this function will be called when we do not have anything placed yet
// we also do not have to check anything below because we have not went to row+1 yet,
// and so we only need to check up from a specific r,c
func canBePlaced(r, c int, board [][]bool) bool {
    n := len(board)
    dirs := [][]int{{-1,0},{-1,-1},{-1,1}}
    for _, dir := range dirs {
        newR := r + dir[0]
        newC := c + dir[1]
        for newR >= 0 && newR < n && newC >= 0 && newC < n {
            if board[newR][newC] == true {
                return false
            }
            // continue in the same direction as long as we are inbound
            newR = newR + dir[0]
            newC = newC + dir[1]
        }  
    }
    return true
} 

