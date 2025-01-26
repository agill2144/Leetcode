func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int)bool{
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] < intervals[j][1]
        }
        return intervals[i][0] < intervals[j][0]
    })
    prev := 0
    count := 0
    for i := 1; i < len(intervals); i++ {
        start, end := intervals[i][0], intervals[i][1]
        prevEnd := intervals[prev][1]
        if start < prevEnd {
            if end < prevEnd {
                // keep curr, drop prev
                prev = i
            }
            // keep prev, drop curr
            count++
        } else {
            // no overlap, move prev ptr to curr i ptr
            prev = i
        }
    }
    return count
}