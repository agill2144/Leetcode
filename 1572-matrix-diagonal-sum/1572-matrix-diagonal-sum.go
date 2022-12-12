func diagonalSum(mat [][]int) int {
    r := 0
    c := 0
    m := len(mat)
    n := len(mat[0])
    sum := 0
    
    for r < m && c < n { // time: m+n
        sum += mat[r][c]
        mat[r][c] *= -1 // mark it visited
        r++
        c++
    }
    
    r = 0
    c = n-1
    
    for r < m && c >= 0 { // time : m+n
        if mat[r][c] > 0 { // only add unvisited nodes
            sum += mat[r][c] 
        }
        r++
        c--
    }
    
    return sum
}