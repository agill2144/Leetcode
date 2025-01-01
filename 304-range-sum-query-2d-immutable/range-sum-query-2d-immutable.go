type NumMatrix struct {
    matrix [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    m := len(matrix)
    n := len(matrix[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            left := 0
            if j-1 >= 0 {left = matrix[i][j-1]}
            top := 0
            if i-1 >= 0 {top = matrix[i-1][j]}
            overlapping := 0
            if i-1 >= 0 && j-1 >= 0 {overlapping = matrix[i-1][j-1]}
            val := left+top+matrix[i][j]-overlapping
            matrix[i][j] = val
        }
    }    
    return NumMatrix{matrix}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    topSum := 0
    if row1-1 >= 0 {topSum = this.matrix[row1-1][col2]}
    leftSum := 0
    if col1-1 >= 0 {leftSum = this.matrix[row2][col1-1]}
    overlapping := 0
    if row1-1 >= 0 && col1-1 >= 0 {overlapping = this.matrix[row1-1][col1-1]}
    return this.matrix[row2][col2] - topSum - leftSum + overlapping
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */