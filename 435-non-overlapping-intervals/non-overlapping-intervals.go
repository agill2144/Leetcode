func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int)bool{
        return intervals[i][0] < intervals[j][0]
    })
    taken := 1
    lastEnd := intervals[0][1]
    for i := 1; i < len(intervals); i++ {
        start, end := intervals[i][0], intervals[i][1]
        if start < lastEnd {
            // overlap
            lastEnd = min(lastEnd, end)
        } else {
            // no overlap
            lastEnd = end
            taken++
        }
    }
    return len(intervals)-taken
}