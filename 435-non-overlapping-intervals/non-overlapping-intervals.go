func eraseOverlapIntervals(intervals [][]int) int {
    if len(intervals) <= 1 {return 0}
    sort.Slice(intervals, func(i, j int)bool{
        return intervals[i][0] < intervals[j][0]
    })
    count := 0
    lastEnd := intervals[0][1]
    for i := 1; i < len(intervals); i++ {
        start, end := intervals[i][0], intervals[i][1]
        if start < lastEnd {
            count++
            if end < lastEnd {lastEnd = end}
        } else {
            lastEnd = end
        }
    }
    return count
}