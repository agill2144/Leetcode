func insert(intervals [][]int, newInterval []int) [][]int {
    if len(intervals) == 0 {return [][]int{newInterval}}
    insertIdx := binarySearch(intervals, newInterval[0])
    tmp := [][]int{}
    tmp = append(tmp, intervals[:insertIdx]...)
    tmp = append(tmp, newInterval)
    tmp = append(tmp, intervals[insertIdx:]...)
    fmt.Println(tmp)
    out := [][]int{tmp[0]}
    for i := 1; i < len(tmp); i++ {
        start, end := tmp[i][0], tmp[i][1]
        prevEnd := out[len(out)-1][1]
        if start <= prevEnd {
            out[len(out)-1][1] = max(prevEnd, end)
        } else {
            out = append(out, tmp[i])
        }
    }
    return out
}

func binarySearch(intervals [][]int, target int) int {
    if target < intervals[0][0] {return 0}
    if target > intervals[len(intervals)-1][0] {return len(intervals)}
    // right most value on left side of target
    left := 0
    right := len(intervals)-1
    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        val := intervals[mid][0]
        if val > target {
            right = mid-1
        } else {
            ans = mid
            left = mid+1
        }
    }
    return ans+1
}