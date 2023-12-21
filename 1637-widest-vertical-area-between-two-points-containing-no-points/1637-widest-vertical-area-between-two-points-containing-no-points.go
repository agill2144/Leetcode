func maxWidthOfVerticalArea(points [][]int) int {
    sort.Slice(points, func(i, j int)bool{
        return points[i][0] < points[j][0]
    })
    maxW := 0
    for i := 1; i < len(points); i++ {
        currX := points[i][0]
        prevX := points[i-1][0]
        if currX-prevX > maxW {maxW = currX-prevX}
    }
    return maxW
}