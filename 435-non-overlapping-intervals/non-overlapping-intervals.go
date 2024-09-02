func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int)bool{
        return intervals[i][0] < intervals[j][0]
    })
    eraseCount := 0
    lastEnd := intervals[0][1]
    for i := 1; i < len(intervals); i++ {
        start, end := intervals[i][0], intervals[i][1]
        if start < lastEnd {
            // overlap, meaning we need to drop this or prev meeting (i.e erase one of them )
            lastEnd = min(lastEnd, end)
            eraseCount++
        } else {
            // no overlap, nothing to drop, nothing to erase
            lastEnd = end
        }
    }
    return eraseCount
}