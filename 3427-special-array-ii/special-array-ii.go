func isArraySpecial(nums []int, queries [][]int) []bool {
    badIdx := []int{}
    n := len(nums)
    for i := 0; i+1 < n; i++ {
        curr := nums[i]
        next := nums[i+1]
        if curr % 2 == 0 && next % 2 == 0 {
            badIdx = append(badIdx, i)
        } else if curr % 2 != 0 && next % 2 != 0 {
            badIdx = append(badIdx, i)
        }   
    }
    out := make([]bool, len(queries))
    for i := 0; i < len(out); i++ {out[i] = true}
    for i := 0; i < len(queries); i++ {
        start, end := queries[i][0], queries[i][1]
        left := 0
        right := len(badIdx)-1
        for left <= right {
            mid := left + (right-left)/2
            val := badIdx[mid]
            if val >= start && val < end {
                out[i] = false
                break
            }
            if val <= start {
                left = mid+1
            } else {
                right=mid-1
            }
        }
    }
    return out
}