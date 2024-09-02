func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int)bool{
        return intervals[i][0] < intervals[j][0]
    })
    taken := 1
    lastEnd := intervals[0][1]
    for i := 1; i < len(intervals); i++ {
        start, end := intervals[i][0], intervals[i][1]
        if start < lastEnd {
            // overlap, meaning we need to drop this or prev meeting (i.e erase one of them )
            // always evaluate the golden thumb rule: "can we take/keep the earliest ending time"
            // yes, take the one that ends the earliest
            // so that 1 big interval like [1,10] ( prev ) does not stop us from picking [1,2][2,3][3,4]
            // the question becomes "should we take [1,10] or [1,2] ? given that the rest of intervals are  [2,3][3,4]" 
            lastEnd = min(lastEnd, end)
        } else {
            // no overlap, nothing to drop, nothing to erase
            lastEnd = end
            taken++
        }
    }
    // erased meeting = n - taken
    return len(intervals) - taken
}