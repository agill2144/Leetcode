func findRightInterval(intervals [][]int) []int {
    out := make([]int, len(intervals))
    idx := map[[2]int]int{}
    for i := 0; i < len(intervals); i++ {
        idx[[2]int{intervals[i][0], intervals[i][1]}] = i
        out[i] = -1
    }
    sort.Slice(intervals, func(i, j int)bool {
        return intervals[i][0] < intervals[j][0]
    })

    // "right interval for an interval i is an interval j such that startJ >= endI and startJ is minimized"
    // find the left-most idx on right side of intervals[i][1]
    out[len(out)-1] = -1
    for i := 0; i < len(intervals); i++ {
        ithOrigIdx := idx[[2]int{intervals[i][0], intervals[i][1]}]
        endI := intervals[i][1]
        // why include ith idx as part of our search?
        // we are looking for right interval of i
        // therefore should'nt this be i+1 ? i.e start from i+1
        // no, because; "Note that i may equal j"
        left := i
        right := len(intervals)-1
        ans := -1
        for left <= right {
            mid := left + (right-left)/2
            startJ :=  intervals[mid][0]
            if startJ >= endI {
                ans = mid
                right = mid-1
            } else {
                left = mid+1
            }
        }
        if ans != -1 {
            origIdx := idx[[2]int{intervals[ans][0], intervals[ans][1]}]
            out[ithOrigIdx] = origIdx
        }
    }
    return out
}