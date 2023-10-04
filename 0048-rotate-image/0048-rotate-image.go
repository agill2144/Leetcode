/*
    approach:
    - transpose and reverse each row
    - transposing a matrix is when a row becomes a col
        - row0 becomes col0
        - row1 becomes col1
        - row2 becomes col2

    - VERY IMPORTANT: even after transposing a matrix
        the diagonal from top left to bottom right ( 0,0 -> m-1,n-1 )
        stays the same original matrix
    
    - How to transpose a matrix?
    - swap values from 1 side of diagonal line(that does not change) with other side of diagonal line
    [1,2,3]        [1,4,7]
    [4,5,6]   ->   [2,5,8]
    [7,8,9]        [3,6,9]    

    The diagonal 1-5-9 stays the same ( even in the transposed matrix )
    swap from one side of diagonal with other
    
    - matrix[0][1] swapped with matrix[1][0]
    - matrix[0][2] swapped with matrix[2][0]
    - matrix[1][2] swapped with matrix[2][1]    
    - swap(matrix[i][j], matrix[j][i])
    
    time = o(n^2) + o(n^2)
    space = o(1)
*/
func rotate(matrix [][]int)  {
    n := len(matrix)
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
    
    for i := 0; i < n; i++ {
        for j := 0; j < n/2; j++ {
            matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
        }
    }
}