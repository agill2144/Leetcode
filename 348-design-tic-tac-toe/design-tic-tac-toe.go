/*
    approach: brute force
    - for each move, check if current player is a winner
        - check top to bottom in the col
        - check left to right in the row
        - check diagonal left to right (0,0 -> n-1,n-1)
        - check diagonal right to left (0,n-1 -> n-1,0)
    
    time = o(n^2)
    space = o(1) if we do not count board to simulate the game
*/


/*
    approach: optimal approach using counters
    - we can maintain a counter for each row for each player
    - we can maintain a counter for each col for each player
    - player1 row counter, player1 col counter
    - player2 row counter, player2 col counter
    - when player1 makes a move at row=1 col=2
        - increment player1 row 1 counter
        - check if player1 has had n markers in this row 1
            - only possible if row 1 counter is now n
        - increment player1 col 2 counter
        - check if player1 has had n markers in col 2
            - only possible if col 2 counter is now n
    - the above handles left -> right , top -> bottom checks
    - does not handle diagonals.

    - player1 will have its own diagonal and anti-diagonal counter (int)
    - player2 will have its own diagonal and anti-diagonal counter (int)
        - anti-diagonal is diagonal from right to left
    - inorder for diagonal path to win, there must be n markers in that diagonal path
    - there are only EVER GOING TO BE 2 paths in diagonal paths, which can contain n markers

    - say n = 4 ( 4x4 board )

    - to identify whether a move is in a diagonal path
        - the row and col will always be the same!
        - r=2;c=2 , r=3;c=3, etc
        - so if a move has row==col, mark its row and col counters, as well as increment diagonal counter

    - to identify whether a move is in a anti-diagonal path
        - row+col will always equal n-1
        - row=0; col=3 ; 0+3 = 3 and n-1 is 3
        - row=1; col=2 ; 1+2 = 3
        - row=2; col=1 ; 2+1 = 3
    
    - anti-diagonal = row+col == n-1
    - diagonal = row==col
    - now we can maintain anti diagonal and diagonal counters for each player 
*/

type stats struct {
    rows []int
    cols []int
    diagonals int
    antiDiagonals int
}


type TicTacToe struct {
    n int
    playerScores map[int]stats
}


func Constructor(n int) TicTacToe {
    return TicTacToe{
        n:n,
        playerScores: map[int]stats{
            1: stats{rows: make([]int, n), cols: make([]int, n)},
            2: stats{rows: make([]int, n), cols: make([]int, n)},
        },
    }
}


func (this *TicTacToe) Move(row int, col int, player int) int {
    pStats := this.playerScores[player]
    // update how many markers this player has in this $row
    pStats.rows[row]++
    // check if this player has filled all this $row ( if row counter == n )
    if pStats.rows[row] == this.n {return player}
    
    // update how many markers this player has in this $col
    pStats.cols[col]++
    // check if this player has filled all this $col ( if col counter == n )
    if pStats.cols[col] == this.n {return player}

    // diagonal ; if r == c 
    if row == col {
        pStats.diagonals++
        // check whether diagonal is full
        if pStats.diagonals == this.n {return player}
    }

    // anti-diagonal; if r+c == n-1
    if row+col == this.n-1 {
        pStats.antiDiagonals++
        // check whether antiDiagonal is full
        if pStats.antiDiagonals == this.n {return player}
    }

    // save updated pStats
    this.playerScores[player] = pStats
    
    // this $player move was not a winner move, therefore return 0
    return 0
}



/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */