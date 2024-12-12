func findRightInterval(intervals [][]int) []int {
    idx := map[[2]int]int{}
    for i := 0; i < len(intervals); i++ {
        idx[[2]int{intervals[i][0], intervals[i][1]}] = i
    }
    sort.Slice(intervals, func(i, j int)bool{
        return intervals[i][0] < intervals[j][0]
    })
    out := make([]int, len(intervals))
    // "right interval for an interval i is an interval j such that startJ >= endI and startJ is minimized"
    // find the left-most idx on right side of intervals[i][1]
    for i := 0; i < len(intervals); i++ {
        ogIdx := idx[[2]int{intervals[i][0], intervals[i][1]}]
        out[ogIdx] = -1
        endI := intervals[i][1]
        // why include ith idx as part of our search?
        // we are looking for right interval of i
        // therefore should'nt this be i+1 ? i.e start from i+1
        // no, because; "Note that i may equal j"
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

/*
    [[3,4],[2,3],[1,2]]
    
    ogIdxs = {
        [3,4]: 0,
        [2,3]: 1,
        [1,2]: 2
    }
    [[1,2], [2,3], [3,4]]
*/
