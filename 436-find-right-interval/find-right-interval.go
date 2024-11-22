func findRightInterval(intervals [][]int) []int {
    idx := map[[2]int]int{}
    for i := 0; i < len(intervals); i++ {
        idx[[2]int{intervals[i][0], intervals[i][1]}] = i
    }
    sort.Slice(intervals, func(i, j int)bool{
        return intervals[i][0] < intervals[j][0]
    })
    out := make([]int, len(intervals))
    for i := 0; i < len(intervals); i++ {
        ogIdx := idx[[2]int{intervals[i][0], intervals[i][1]}]
        out[ogIdx] = -1
        endI := intervals[i][1]
        left := i
        right := len(intervals)-1
        for left <= right {
            mid := left + (right-left)/2
            startJ := intervals[mid][0]
            if startJ >= endI {
                out[ogIdx] = idx[[2]int{intervals[mid][0], intervals[mid][1]}]
                right = mid-1
            } else {
                left = mid+1
            }
        }
    }
    return out
}