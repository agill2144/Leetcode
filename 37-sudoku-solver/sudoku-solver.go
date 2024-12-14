type Solution struct {
    N                    int
    n                    int
    rows, columns, boxes [][]int
    board                [][]byte
    sudokuSolved         bool
}

func Constructor() Solution {
    n := 3
    N := n * n
    rows := make([][]int, N)
    columns := make([][]int, N)
    boxes := make([][]int, N)
    for i := range rows {
        rows[i] = make([]int, N+1)
        columns[i] = make([]int, N+1)
        boxes[i] = make([]int, N+1)
    }
    return Solution{
        N:            N,
        n:            n,
        rows:         rows,
        columns:      columns,
        boxes:        boxes,
        sudokuSolved: false,
    }
}

func (this *Solution) couldPlace(d, row, col int) bool {
    idx := (row/this.n)*this.n + col/this.n
    return this.rows[row][d]+this.columns[col][d]+this.boxes[idx][d] == 0
}

func (this *Solution) placeNumber(d, row, col int) {
    idx := (row/this.n)*this.n + col/this.n
    this.rows[row][d]++
    this.columns[col][d]++
    this.boxes[idx][d]++
    this.board[row][col] = byte(d) + '0'
}

func (this *Solution) removeNumber(d, row, col int) {
    idx := (row/this.n)*this.n + col/this.n
    this.rows[row][d]--
    this.columns[col][d]--
    this.boxes[idx][d]--
    this.board[row][col] = '.'
}

func (this *Solution) placeNextNumbers(row, col int) {
    if (col == this.N-1) && (row == this.N-1) {
        this.sudokuSolved = true
    } else {
        if col == this.N-1 {
            this.backTrack(row+1, 0)
        } else {
            this.backTrack(row, col+1)
        }
    }
}

func (this *Solution) backTrack(row, col int) {
    if this.board[row][col] == '.' {
        for d := 1; d < 10; d++ {
            if this.couldPlace(d, row, col) {
                this.placeNumber(d, row, col)
                this.placeNextNumbers(row, col)
                if !this.sudokuSolved {
                    this.removeNumber(d, row, col)
                }
            }
        }
    } else {
        this.placeNextNumbers(row, col)
    }
}

func solveSudoku(board [][]byte) {
    s := Constructor()
    s.board = board
    for i := 0; i < s.N; i++ {
        for j := 0; j < s.N; j++ {
            num := board[i][j]
            if num != '.' {
                d := int(num - '0')
                s.placeNumber(d, i, j)
            }
        }
    }
    s.backTrack(0, 0)
    board = s.board
}