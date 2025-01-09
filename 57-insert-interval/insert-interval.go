/*
    approach: 
    - search, insert, merge
    - given sorted intervals, use binary search to find the correct insert idx position
        - "find right most val on left side of target"
    - create a tmp array, and insert newInterval as well as old ones
        - there may be overalap
    - now finally, merge overlapping intervals
        - overlapping in this question is; start <= prevEnd time
    
    time;
    - search = o(logn)
    - insert = o(n)
    - merge = o(n)
    total time = o(logn) + o(2n)

    space = o(n) for tmp array
*/
func insert(intervals [][]int, newInterval []int) [][]int {
    if len(intervals) == 0 {return [][]int{newInterval}}
    idx := search(intervals, newInterval[0])
    // space = o(n)
    tmp := [][]int{}
    tmp = append(tmp, intervals[:idx]...)
    tmp = append(tmp, newInterval)
    tmp = append(tmp, intervals[idx:]...)

    // space = o(n) but will be o(1) if output array is not counted
    // some people count it, some dont, i dont
    // therefore o(1) space for this merged array
    merged := [][]int{tmp[0]}
    for i := 1; i < len(tmp); i++ {
        start, end := tmp[i][0], tmp[i][1]
        prevStart, prevEnd := merged[len(merged)-1][0],merged[len(merged)-1][1]
        if start <= prevEnd {
            merged[len(merged)-1][0] = min(start, prevStart)
            merged[len(merged)-1][1] = max(end, prevEnd)
        } else {
            merged = append(merged, tmp[i])
        }
    }
    return merged

}


// time = o(logn)
// space = o(1)
func search(intervals [][]int, target int) int {
    if target < intervals[0][0] {return 0}
    if target > intervals[len(intervals)-1][0] {return len(intervals)}
    left := 0
    right := len(intervals)-1
    // find right most val on left side of target
    ans := 0
    for left <= right {
        mid := left+(right-left)/2
        val := intervals[mid][0]
        if val <= target {
            ans = mid
            left = mid+1 // since we want the right most value on left side of target
        } else {
            right = mid-1
        }
    }
    // why +1 ?
    // because this function is supposed to return the insert idx position
    return ans+1
}