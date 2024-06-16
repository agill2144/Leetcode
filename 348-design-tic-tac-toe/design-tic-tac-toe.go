type TicTacToe struct {
    n int
    rows []int
    cols []int
    diaLeft int
    diaRight int

}


func Constructor(n int) TicTacToe {
    return TicTacToe{
        n:n,
        rows: make([]int, n),
        cols: make([]int, n),
    }
}


func (this *TicTacToe) Move(row int, col int, player int) int {
    count := 1
    if player == 2 {count = -1}

    this.rows[row] += count
    if abs(this.rows[row]) == this.n {return player}

    this.cols[col] += count
    if abs(this.cols[col]) == this.n {return player}

    if row-col == 0 {
        this.diaLeft += count
        if abs(this.diaLeft) == this.n {return player}
    } 

    if row+col == this.n-1 {
        this.diaRight += count
        if abs(this.diaRight) == this.n {return player}
    }

    return 0
}


func abs(x int) int {
    if x < 0 {return x*-1}
    return x
}

/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */