func insert(intervals [][]int, newInterval []int) [][]int {
    if len(intervals) == 0 {
        return [][]int{newInterval}
    }

    left := 0
    right := len(intervals)-1
    idx := -1
    for left <= right {
        mid := left + (right-left)/2
        val := intervals[mid][0]
        valOnLeft := val <= newInterval[0]
        if valOnLeft {
            idx = mid
            left = mid+1
        } else {
            right = mid-1
        }
    }
    idx++
    // case1: no valid idx found
    // case2: valid idx found

    tmp := [][]int{}
    i := 0
    for i < idx {tmp = append(tmp, intervals[i]);i++}
    tmp = append(tmp, newInterval)
    for i < len(intervals) {tmp = append(tmp, intervals[i]); i++}
    fmt.Println(idx,tmp)

    // merge intervals code 
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