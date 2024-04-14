func minRectanglesToCoverPoints(points [][]int, w int) int {
    sort.Slice(points, func(i, j int)bool{
        return points[i][0] < points[j][0]
    })
    count := 0
    i := 0
    for i < len(points) {
        upToIncluding := points[i][0] + w
        for i < len(points) && points[i][0] <= upToIncluding {
            i++
        }
        count++
    }
    return count
}