func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int)bool{
        return intervals[i][0] < intervals[j][0]
    })
    lastValidIdx := 0
    countRemoved := 0
    for i := 1; i < len(intervals); i++ {
        currStart := intervals[i][0]
        prevEnd := intervals[lastValidIdx][1]
        if currStart < prevEnd {
            if intervals[i][1] < intervals[lastValidIdx][1] {lastValidIdx=i}
            countRemoved++
        } else {
            lastValidIdx = i
        }
    }
    return countRemoved
}