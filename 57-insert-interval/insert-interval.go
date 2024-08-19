func insert(intervals [][]int, newInterval []int) [][]int {
    n := len(intervals)
    if n == 0 {return [][]int{newInterval}}
    idx := binarySearch(intervals, newInterval[0])
    
    tmp := [][]int{}
    i := 0
    for i < idx {tmp = append(tmp, intervals[i]);i++}
    tmp = append(tmp, newInterval)
    for i < len(intervals) {tmp = append(tmp, intervals[i]); i++}
    
    out := [][]int{}
    for i := 0; i < len(tmp); i++ {
        if len(out) != 0 && tmp[i][0] <= out[len(out)-1][1] {
            out[len(out)-1][0] = min(tmp[i][0], out[len(out)-1][0])
            out[len(out)-1][1] = max(tmp[i][1], out[len(out)-1][1])
        } else {
            out = append(out, tmp[i])
        }
    }
    return out
}

func binarySearch(intervals [][]int, target int) int {
    // 1,3,6,9,10 ; target 2 -> return idx 1 ( rightmost number on left side of target )
    // 6,7,8,9; target 2 -> return idx 0
    // 1,2,3; target 5 -> return 2

    left := 0
    right := len(intervals)-1
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        start := intervals[mid][0]
        if start <= target {
            ans = mid
            // keep searching right because we want the rightmost value on the left side of target
            left = mid+1
        }  else {
            right = mid-1
        }
    }

    // why +1 ?
    // because with above binary search, we found the rightmost idx on left side target
    // but we need to insert the new interval just after this idx position
    // therefore +1 ; to indicate that this is the idx position to insert at
    return ans+1
}  