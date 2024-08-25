func insert(intervals [][]int, newInterval []int) [][]int {
    idx := binarySearch(intervals, newInterval[0])
    tmp := [][]int{}
    tmp = append(tmp, intervals[:idx]...)
    tmp = append(tmp, newInterval)
    tmp = append(tmp, intervals[idx:]...)
    merged := [][]int{tmp[0]}
    for i := 1; i < len(tmp); i++ {
        start, end := tmp[i][0], tmp[i][1]
        prevStart, prevEnd := merged[len(merged)-1][0], merged[len(merged)-1][1]
        if start <= prevEnd {
            merged[len(merged)-1][0] = min(start, prevStart)
            merged[len(merged)-1][1] = max(end, prevEnd)
        } else {
            merged = append(merged, tmp[i])
        }
    }
    return merged
}

func binarySearch(intervals [][]int, target int) int {
    n := len(intervals)
    if n == 0 || target < intervals[0][0] {
        return 0
    }
    left := 0
    right := n-1
    idx := -1
    for left <= right {
        mid := left + (right-left)/2
        val := intervals[mid][0]
        if val < target {
            idx = mid
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return idx+1
}