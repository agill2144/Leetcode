func generate(numRows int) [][]int {
    matrix := [][]int{{1}}
    for i := 1; i < numRows; i++ {
        prevRow := matrix[i-1]
        row := []int{1}
        
        for j := 1; j < i; j++ {
            row = append(row, prevRow[j-1]+prevRow[j])
        }
        row = append(row, 1)
        matrix = append(matrix, row)
    }
    return matrix
}