func merge(intervals [][]int) [][]int {
    if len(intervals) <= 1 {return intervals}

    sort.SliceStable(intervals, func(i, j int)bool{
        return intervals[i][0] < intervals[j][0]
    })
    
    merged := [][]int{intervals[0]}
    
    for i := 1; i < len(intervals); i++ {
        if intervals[i][0] <= merged[len(merged)-1][1] {
            merged[len(merged)-1][0] = min(intervals[i][0], merged[len(merged)-1][0])
            merged[len(merged)-1][1] = max(intervals[i][1], merged[len(merged)-1][1])
        } else {
            merged = append(merged, intervals[i])
        }
    }
    return merged
}

func min(x, y int) int {
    if x < y {return x}
    return y
}

func max(x, y int) int {
    if x > y {return x}
    return y
}