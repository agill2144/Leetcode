func insert(intervals [][]int, newInterval []int) [][]int {
    if newInterval == nil || len(newInterval) != 2  {
        return intervals
    }
    
    intervals = append(intervals, newInterval)
    for i := len(intervals)-2; i >= 0; i-- {
        next := intervals[i+1]
        curr := intervals[i]
        if next[0] < curr[0] {
            intervals[i+1], intervals[i] = intervals[i], intervals[i+1]
        }
    }
    
    out := [][]int{}
    out = append(out, intervals[0])
    for i := 1; i < len(intervals); i++ {
        curr := intervals[i]
        prev := out[len(out)-1]
        if curr[0] <= prev[1] {
            out = out[:len(out)-1]
            newInterval := []int{min(prev[0], curr[0]), max(prev[1],curr[1])}
            out = append(out, newInterval)
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