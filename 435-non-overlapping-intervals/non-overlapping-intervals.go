func eraseOverlapIntervals(intervals [][]int) int {
    // [[1,2],[1,3],[2,3],[3,4]]   
    sort.Slice(intervals, func(i, j int)bool {
        return intervals[i][0] < intervals[j][0]
    })
    count := 0
    prev := 0
    for i := 1; i < len(intervals); i++ {
        start, end := intervals[i][0], intervals[i][1]
        prevEnd := intervals[prev][1]
        if start < prevEnd {
            count++
            if end < prevEnd {
                prev = i
            }
        } else {
            prev = i
        }
    }
    return count
}