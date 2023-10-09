func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int)bool {
        return intervals[i][0] < intervals[j][0]
    })
    out := [][]int{intervals[0]}
    for i := 1; i < len(intervals); i++ {
        prev := out[len(out)-1]
        curr := intervals[i]
        if curr[0] <= prev[1] {
            out[len(out)-1][0] = min(prev[0], curr[0])
            out[len(out)-1][1] = max(prev[1], curr[1])
        } else {
            out = append(out, curr)
        }
    }
    return out
}

func min(x, y int) int {
    if x < y {return x}
    return y
}
func max(x, y int) int {
    if x > y {return x}
    return y
}