func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int)bool{
        return intervals[i][0] < intervals[j][0]
    })
    out := [][]int{intervals[0]}
    for i := 1; i < len(intervals); i++ {
        start := intervals[i][0]
        end := intervals[i][1]
        prevStart := out[len(out)-1][0]
        prevEnd := out[len(out)-1][1]
        
        if start <= prevEnd {
            out[len(out)-1][0] = min(start, prevStart)
            out[len(out)-1][1] = max(end, prevEnd)
        } else {
            out = append(out, intervals[i])
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