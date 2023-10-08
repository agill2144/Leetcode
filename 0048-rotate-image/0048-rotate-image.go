/*
    approach: brute force
    - create a new matrix
    - map input i,j -> new matrix j,(m-i-1)
    - copy new matrix values to input matrix
    
    time = o(mn) + o(mn)
    space = o(mn)
*/
func rotate(matrix [][]int)  {
    n := len(matrix)
    tmp := make([][]int, n)
    for i := 0; i < n; i++ {
        tmp[i] = make([]int, n)
    }
    
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            tmp[j][n-1-i] = matrix[i][j]
        }
    }
     for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            matrix[i][j] = tmp[i][j]
        }
     }
}