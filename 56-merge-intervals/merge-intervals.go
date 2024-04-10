func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int)bool{
        return intervals[i][0] < intervals[j][0]
    })
    out := [][]int{intervals[0]}
    for i := 1; i < len(intervals); i++ {
        curr := intervals[i]
        prev := out[len(out)-1]
        if curr[0] <= prev[1] {
            out[len(out)-1][0] = min(curr[0], prev[0])
            out[len(out)-1][1] = max(curr[1], prev[1])
        } else {
            out = append(out, curr)
        }
    }
    return out
}