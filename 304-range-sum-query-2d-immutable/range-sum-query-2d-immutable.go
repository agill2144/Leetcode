type NumMatrix struct {
    fullSum [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    m := len(matrix)
    n := len(matrix[0])
    fullSum := make([][]int, m)
    for i := 0; i < m; i++ {
        fullSum[i] = make([]int, n)
        for j := 0; j < n; j++ {
            leftSum := 0
            if j-1 >= 0 {leftSum = fullSum[i][j-1]}
            topSum := 0
            if i-1 >= 0 {topSum = fullSum[i-1][j]}
            overlap := 0
            if i-1 >= 0 && j-1 >= 0 {
                overlap = fullSum[i-1][j-1]
            }
            fullSum[i][j] = matrix[i][j] + leftSum + topSum - overlap
        }
    }
    return NumMatrix{fullSum}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    total := this.fullSum[row2][col2]
    if row1-1 >= 0 {
        total -= this.fullSum[row1-1][col2]
    }
    if col1-1 >= 0 {
        total -= this.fullSum[row2][col1-1]
    }
    if row1-1 >= 0 && col1-1 >= 0 {
        total += this.fullSum[row1-1][col1-1]
    }
    return total
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */