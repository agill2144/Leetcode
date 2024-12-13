type NumMatrix struct {
    prefixSum [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    m := len(matrix)
    n := len(matrix[0])
    prefixSum := make([][]int, m)
    for i := 0; i < m; i++ {
        prefixSum[i] = make([]int, n)
        for j := 0; j < n; j++ {
            prev := 0
            if j-1 >= 0 {prev = prefixSum[i][j-1]}
            above := 0
            if i-1 >= 0 {above = prefixSum[i-1][j]}
            overlap := 0
            if i-1 >= 0 && j-1 >= 0 {overlap = prefixSum[i-1][j-1]}
            prefixSum[i][j] = matrix[i][j] + prev + above - overlap
        }
    }
    return NumMatrix{prefixSum}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    // right corner sum 
    total := this.prefixSum[row2][col2]
    // remove left panel sum
    if col1-1 >= 0 {
        total -= this.prefixSum[row2][col1-1]
    }
    // remove top panel sum
    if row1-1 >= 0 {
        total -= this.prefixSum[row1-1][col2]
    }
    // add back the overlapping cells that got removed twice
    if row1-1 >= 0 && col1-1 >= 0 {
        total += this.prefixSum[row1-1][col1-1]
    }
    return total

}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */