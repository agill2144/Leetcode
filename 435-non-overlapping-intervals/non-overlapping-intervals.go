func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int)bool {
        return intervals[i][1] < intervals[j][1]
    })
    lastEnd := math.MinInt64
    count := 0 // count taken
    for i := 0; i < len(intervals); i++ {
        if lastEnd == math.MinInt64 || intervals[i][0] >= lastEnd{
            lastEnd = intervals[i][1]
            count++
        }
    }
    return len(intervals)-count
}