func findMinArrowShots(points [][]int) int {
    sort.Slice(points, func(i, j int)bool{
        if points[i][0] == points[j][0] {
            return points[i][1] < points[j][1]
        }
        return points[i][0] < points[j][0]
    })
    count := 1
    groupEnd := 0
    for i := 1; i < len(points); i++ {
        start, end := points[i][0], points[i][1]
        if start <= points[groupEnd][1] {
            if end < points[groupEnd][1] {
                groupEnd = i
            }
        } else {
            count++
            groupEnd = i
        }
    }
    return count
}