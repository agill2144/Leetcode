// what a stupid problem to waste time on ... fuck you google
func numMagicSquaresInside(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    if m < 3 || n < 3 {return 0}
    count := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i+2 < m && j+2 < n {
                if isValid(i, j, grid) {
                    count++
                }
            }
        }
    }
    return count
}


func isValid(r, c int, grid [][]int) bool {
    rowSum := -1 // not yet identified
    colSum := -1 // not yet identified
    set := map[int]struct{}{}
    for i := r; i <= r+2; i++ { // o(3)
        rSum := 0
        for j := c; j <= c+2; j++ { // o(3)
            if _, ok := set[grid[i][j]]; ok {return false}
            if grid[i][j] < 1 || grid[i][j] > 9 {return false}
            rSum += grid[i][j]
            set[grid[i][j]] = struct{}{}
        }
        if rowSum == -1 {rowSum = rSum}
        if rSum != rowSum {return false}
    }

    for j := c; j <= c+2; j++ { // o(3)
        cSum := 0
        for i := r; i <= r+2; i++ { // o(3)
            cSum += grid[i][j]
        }      
        if colSum == -1 {colSum = cSum}
        if cSum != colSum {return false}
        if colSum != rowSum {return false}  
    }

    lrSum := 0
    lrr := r; lrc := c
    for lrr <= r+2 && lrc <= c+2 { // o(3)
        lrSum += grid[lrr][lrc]
        lrr++
        lrc++
    }

    rlSum := 0
    rlr := r; rlc := c+2
    for rlr <= r+2 && rlc <= c+2 { // o(3)
        rlSum += grid[rlr][rlc]
        rlr++
        rlc--
    }
    return lrSum == rlSum && rlSum == rowSum 
}