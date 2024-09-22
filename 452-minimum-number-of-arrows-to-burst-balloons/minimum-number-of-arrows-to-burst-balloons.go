func findMinArrowShots(points [][]int) int {
    sort.Slice(points, func(i, j int)bool{
        return points[i][0] < points[j][0]
    })
    count := 1
    lastEnd := points[0][1]
    for i := 0; i < len(points); i++ {
        start, end := points[i][0], points[i][1]
        // is it overlapping?
        if start <= lastEnd {
            lastEnd = min(lastEnd, end)
        } else {
            // no overlap, 
            // need a new arrow
            count++
            lastEnd = end
        }
    }
    return count
}