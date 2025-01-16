func insert(intervals [][]int, newInterval []int) [][]int {
    if len(intervals) == 0 {return [][]int{newInterval}}
    // move i ptr where this interval should be inserted at
    out := [][]int{}
    i := 0
    for i < len(intervals) {
        if intervals[i][0] <= newInterval[0] {
            out = append(out, intervals[i])
            i++
            continue
        }
        break
    }

    // now insert the newInterval
    // but could this new interval overlap with last interval inserted?
    // yes
    start, end := newInterval[0], newInterval[1]
    if len(out) > 0 && start <= out[len(out)-1][1] {
        out[len(out)-1][1] = max(out[len(out)-1][1],end)
    } else {
        out = append(out, newInterval)
    }
    // now add the rest of remaining intervals
    // while merging overlaps
    for i < len(intervals) {
        start, end := intervals[i][0], intervals[i][1]
        lastStart, lastEnd := out[len(out)-1][0], out[len(out)-1][1]
        if start <= lastEnd {
            out[len(out)-1][0] = min(lastStart, start)
            out[len(out)-1][1] = max(lastEnd,end)
        } else {
            out = append(out, intervals[i])
        }
        i++
    }
    return out

}