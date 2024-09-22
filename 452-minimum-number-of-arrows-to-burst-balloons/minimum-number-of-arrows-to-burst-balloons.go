func findMinArrowShots(points [][]int) int {
    sort.Slice(points, func(i, j int)bool{
        return points[i][1] < points[j][1]
    })
    count := 1
    lastEnd := points[0][1]
    for i := 0; i < len(points); i++ {
        start, end := points[i][0], points[i][1]
        // is it overlapping?
        if start <= lastEnd {
            // if yes, the last arrow we shot most likely destroyed this ballon
            // since this ballon overlapped in x-axis with the last
            // continue with same arrow
            continue
        } else {
            // need a new arrow
            count++
            lastEnd = end
        }
    }
    return count
}