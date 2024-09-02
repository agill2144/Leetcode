func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int)bool{
        return intervals[i][0] < intervals[j][0]
    })
    n := len(intervals)
    countTaken := 1
    prevEnd := intervals[0][1]
    for i := 1; i < len(intervals); i++ {
        start, end := intervals[i][0], intervals[i][1]
        if start < prevEnd {
            prevEnd = min(prevEnd, end)
        } else {
            prevEnd = end
            countTaken++
        }
    }
    return n-countTaken
}