func insert(intervals [][]int, newInterval []int) [][]int {
    n := len(intervals)
    if n == 0 {return [][]int{newInterval}}
    out := [][]int{}
    out = append(out, intervals[0])
    i := 1
    used := false

    if newInterval[0] < intervals[0][0] {
        out[len(out)-1] = newInterval
        used = true
        i = 0
    }
    for i < n || !used {
        start := math.MaxInt64
        end := math.MaxInt64
        if i < n {start = intervals[i][0]; end = intervals[i][1]}
        newStart := newInterval[0]
        newEnd := newInterval[1]
        lastStart := out[len(out)-1][0]
        lastEnd := out[len(out)-1][1]

        if newStart < start && !used {
            if newStart <= lastEnd {
                out[len(out)-1][0] = min(lastStart, newStart)
                out[len(out)-1][1] = max(lastEnd, newEnd)
            } else {
                out = append(out, newInterval)
            }
            used = true
        } else {
            if start <= lastEnd {
                out[len(out)-1][0] = min(lastStart, start)
                out[len(out)-1][1] = max(lastEnd, end)
            } else {
                out = append(out, intervals[i])   
            }
            i++
        }
    }
    return out
}