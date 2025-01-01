type NumMatrix struct {
    matrix [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    m := len(matrix)
    n := len(matrix[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if j-1 >= 0 {matrix[i][j] += matrix[i][j-1]}
            if i-1 >= 0 {matrix[i][j] += matrix[i-1][j]}
            // remove the overlapping sum that was added twice
            // once it was added from left cell and other time it was added from top cell
            if i-1 >= 0 && j-1 >= 0 {matrix[i][j] -= matrix[i-1][j-1]}
        }
    }
    return NumMatrix{matrix}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    fullSum := this.matrix[row2][col2]
    // remove top row
    if row1-1 >= 0 {fullSum -= this.matrix[row1-1][col2]}
    // remove left col
    if col1-1 >= 0 {fullSum -= this.matrix[row2][col1-1]}
    // add the overlapping sum that was removed twice
    if row1-1 >= 0 && col1-1 >= 0 {fullSum += this.matrix[row1-1][col1-1]}
    return fullSum
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */