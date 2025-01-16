func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int)bool {
        return intervals[i][0] < intervals[j][0]
    })
    lastKept := 0
    count := 0
    for i := 1; i < len(intervals); i++ {
        start := intervals[i][0]
        prevEnd := intervals[lastKept][1]
        if start < prevEnd {
            count++
            // keep the one that ends earliest, that way we have reduced chances of overlapping
            // keep curr instead of lastKept when curr ends before lastKept ends
            if intervals[i][1] < prevEnd {lastKept = i}
        } else {
            lastKept = i
        }
    }
    return count

}