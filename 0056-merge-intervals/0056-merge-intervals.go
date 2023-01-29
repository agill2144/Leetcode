func merge(intervals [][]int) [][]int {
    sort.SliceStable(intervals, func(i, j int)bool{
        return intervals[i][0] < intervals[j][0]
    })
    out := [][]int{intervals[0]}
    for i := 1; i < len(intervals); i++ {
        start := intervals[i][0]
        end := intervals[i][1]
        
        prev := out[len(out)-1]
        prevStart := prev[0]
        prevEnd := prev[1]
        
        if start > prevEnd {
            out = append(out, []int{start, end})
        } else {
            newEnd := max(end, prevEnd)
            newStart := min(start, prevStart)
            out = out[:len(out)-1]
            out = append(out, []int{newStart, newEnd})
        }
    }
    return out
    
}

func max(x, y int) int {
    if x > y {return x}
    return y
}

func min(x, y int) int {
    if x < y {return x}
    return y
}