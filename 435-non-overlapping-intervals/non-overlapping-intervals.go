func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int)bool {
        return intervals[i][0] < intervals[j][0]
    })
    lastKept := 0
    count := 0
    for i := 1; i < len(intervals); i++ {
        start, end := intervals[i][0], intervals[i][1]
        prevEnd := intervals[lastKept][1]
        if start < prevEnd {
            count++
            if end < prevEnd {lastKept = i}
        } else {
            lastKept = i
        }
    }
    return count

}