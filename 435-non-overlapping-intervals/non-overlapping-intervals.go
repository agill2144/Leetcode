func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int)bool{
        return intervals[i][0] < intervals[j][0]
    })
    n := len(intervals)
    prev := 0
    taken := 1
    for i := 1; i < n; i++ {
        start, end := intervals[i][0], intervals[i][1]
        if start < intervals[prev][1] {
            if end < intervals[prev][1] {
                // keep curr, drop prev
                prev = i
            }
            // otherwise, we are dropping current and keeping prev
        } else {
            // no overlap, meaning we can take current interval
            prev = i
            taken++
        }
    }
    return n-taken
}