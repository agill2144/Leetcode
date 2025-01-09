/*
    approach: brute force
    - append
    - sort
    - merge overlapping
*/
func insert(intervals [][]int, newInterval []int) [][]int {
    if newInterval == nil || len(newInterval) == 0 {return intervals}
    
    // append
    intervals = append(intervals, newInterval)
    
    // sort
    sort.Slice(intervals, func(i, j int)bool{
        return intervals[i][0] < intervals[j][0]
    })

    // merge overlapping
    out := [][]int{intervals[0]}
    for i := 1; i < len(intervals); i++ {
        start, end := intervals[i][0], intervals[i][1]
        prevEnd := out[len(out)-1][1]
        if start <= prevEnd {
            out[len(out)-1][0] = min(start, out[len(out)-1][0])
            out[len(out)-1][1] = max(end, out[len(out)-1][1])
        } else {
            out = append(out, intervals[i])
        }
    }
    return out
}