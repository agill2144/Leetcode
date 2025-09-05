func findRightInterval(intervals [][]int) []int {
    idx := map[[2]int]int{}
    for i := 0; i < len(intervals); i++ {
        idx[[2]int{intervals[i][0], intervals[i][1]}] = i
    }

    sort.Slice(intervals, func(i, j int)bool{
        return intervals[i][0] < intervals[j][0]
    })

    res := make([]int, len(intervals))
    for i := 0; i < len(intervals); i++ {
        // The right interval for an interval i is an interval j 
        // such that startj >= endi and startj is minimized. Note that i may equal j
        // so each endi we need leftmost startj on the right side of endi
        ogIdx := idx[[2]int{intervals[i][0], intervals[i][1]}]
        res[ogIdx] = -1
        left := i
        right := len(intervals)-1
        endi := intervals[i][1]
        for left <= right {
            mid := left + (right-left)/2
            startj := intervals[mid][0]
            if startj >= endi {
                res[ogIdx] = idx[[2]int{intervals[mid][0], intervals[mid][1]}]
                right = mid-1                
            } else {
                left = mid+1
            }
        }
    }
    return res
}