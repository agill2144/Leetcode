func generate(numRows int) [][]int {
    out := make([][]int, numRows)
    for i := 0; i < numRows; i++ {
        desiredSize := i+1
        out[i] = make([]int, desiredSize)
        out[i][0] = 1
        out[i][desiredSize-1] = 1
        for j := 1; j < desiredSize-1; j++ {
            
            out[i][j] = out[i-1][j]+out[i-1][j-1]
        }
    }
    return out
}