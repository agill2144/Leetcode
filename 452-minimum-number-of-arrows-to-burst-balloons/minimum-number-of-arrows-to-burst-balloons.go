func findMinArrowShots(points [][]int) int {
    sort.Slice(points, func(i, j int)bool{
        return points[i][0] < points[j][0]
    })
    count := 1
    prevEnd := points[0][1]
    for i := 1; i < len(points); i++ {
        start, end := points[i][0], points[i][1]
        if start <= prevEnd {
            prevEnd = min(end,prevEnd)
        } else {
            count++
            prevEnd = points[i][1]
        }
    }
    return count
}

/*
    [10,16],[2,8],[1,6],[7,12]]

    [1,6],[2,8],[7,12],[10,16]


*/