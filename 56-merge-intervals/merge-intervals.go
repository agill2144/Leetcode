func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int)bool{
        return intervals[i][0] < intervals[j][0]
    })
    out := [][]int{intervals[0]}
    for i := 1; i < len(intervals); i++ {
        start, end := intervals[i][0], intervals[i][1]
        prevEnd := out[len(out)-1][1]
        if start <= prevEnd {
            out[len(out)-1][1] = max(end, prevEnd)
        } else {
            out = append(out, intervals[i])
        }
    }
    return out
}