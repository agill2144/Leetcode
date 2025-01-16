/*
    approach: 2 passes
    1. insert whatever we can from intervals list first
        - any interval whose start time is before-or-equal-to newIntervals
        - start time, it can be placed before newInterval
    2. now add the new interval
        - its possible that this new interval
        - could overlap with what we last inserted
        - if yes, merge the overlap
        - if no, insert new interval as-is
    3. now add the remaining left over intervals
        - while merging if there 
            are overlap with last inserted interval

    tc = o(n)
    sc = o(1)
*/
func insert(intervals [][]int, newInterval []int) [][]int {
    if len(intervals) == 0 {return [][]int{newInterval}}
    out := [][]int{}
    i := 0
    n := len(intervals)

    // insert whatever we can from intervals list first
    // any interval whose start time is before-or-equal-to newIntervals
    // start time, it can be placed before newInterval
    for i < n && intervals[i][0] <= newInterval[0] {
        out = append(out, intervals[i])
        i++
    }

    // now place the new interval
    // NOTE: its possible that this new interval
    // could overlap with what we last inserted
    // if yes, merge the overlap
    // if no, insert new interval as-is
    if len(out) > 0 && newInterval[0] <= out[len(out)-1][1] {
        out[len(out)-1][1] = max(out[len(out)-1][1], newInterval[1])
    } else {
        out = append(out, newInterval)
    }

    // now add the remaining left over intervals
    // while merging if there are overlap with last inserted interval
    for i < n {
        start, end := intervals[i][0], intervals[i][1]
        prevStart, prevEnd := out[len(out)-1][0], out[len(out)-1][1]
        if start <= prevEnd {
            out[len(out)-1][0] = min(start, prevStart)
            out[len(out)-1][1] = max(end, prevEnd)
        } else {
            out = append(out, intervals[i])
        }
        i++
    }
    return out

    
}