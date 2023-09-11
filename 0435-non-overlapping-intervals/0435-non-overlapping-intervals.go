func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] < intervals[j][1]
        } 
        return intervals[i][0] < intervals[j][0]
    })
    // count all the overlapping intervals - and store the overlapping ones as dirty
    // i.e pretend like they are deleted already
    count := 0
    prev := 0
    for i := 1; i < len(intervals); i++ {
        
        if intervals[i][0] < intervals[prev][1] {
            // be greedy and pick the shorter time between curr and prev
            // i.e if this interval end time is smaller, then use this interval for next prev ( i.e delete prev interval )
            // otherwise delete this interval and keep prev as is
            if intervals[i][1] < intervals[prev][1] {prev = i}
            count++
        } else {
            prev = i
        }
    }
    return count
}