func areaOfMaxDiagonal(dimensions [][]int) int {
    maxArea := 0
    var maxDiag float64
    for i := 0; i < len(dimensions); i++ {
        l := dimensions[i][0]
        w := dimensions[i][1]
        diag := math.Sqrt(float64(l*l) + float64(w*w))
        area := l*w
        if diag > maxDiag {
            maxDiag = diag
            maxArea = area
        } else if diag == maxDiag {
            maxArea = max(maxArea, area)
        }
    }
    return maxArea
}