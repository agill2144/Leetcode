func generate(numRows int) [][]int {
    out := make([][]int, numRows)
    for i := 0; i < numRows; i++ {
        desiredSize := i+1
        out[i] = make([]int, desiredSize)
        out[i][0] = 1
        out[i][desiredSize-1] = 1
        for j := 1; j < desiredSize-1; j++ {
            up := 0
            upLeft := 0
            if i-1 >= 0 {
                up = out[i-1][j]
            }
            if j-1 >= 0 {
                upLeft = out[i-1][j-1]
            }
            out[i][j] = up+upLeft
        }
    }
    return out
}