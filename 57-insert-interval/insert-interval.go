func insert(intervals [][]int, newInterval []int) [][]int {
    if len(newInterval) == 0 {return intervals}
    if len(intervals) == 0 {return [][]int{newInterval}}
    i := 0
    n := len(intervals)
    out := [][]int{}
    for i < n && intervals[i][1] < newInterval[0] {
        out = append(out, intervals[i])
        i++
    }
    for i < n && intervals[i][0] <= newInterval[1] {
        newInterval[0] = min(newInterval[0], intervals[i][0])
        newInterval[1] = max(newInterval[1], intervals[i][1])
        i++
    }
    out = append(out, newInterval)
    for i < n {
        out = append(out, intervals[i])
        i++
    }
    return out
}