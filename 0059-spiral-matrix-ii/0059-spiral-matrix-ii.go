func generateMatrix(n int) [][]int {
    out := make([][]int, n)
    for i, _ := range out {
        out[i] = make([]int, n)
    }
    
    x := 1
    top := 0
    left := 0
    right := n-1
    bottom := n-1
    
    for left <= right && top <= bottom {
        
        for i := left; i <= right; i++ {
            out[top][i] = x
            x++
        }
        top++
        
        if left <= right && top <= bottom {
            for i := top; i <= bottom; i++ {
                out[i][right] = x
                x++
            }
        }
        right--
        
        if left <= right && top <= bottom {
            for i := right; i >= left; i-- {
                out[bottom][i] = x
                x++
            }
        }
        bottom--
        
        if left <= right && top <= bottom {
            for i := bottom; i >= top; i-- {
                out[i][left] = x
                x++
            }
        }
        left++
    }
    return out
}