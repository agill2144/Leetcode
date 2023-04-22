func findMissingRanges(nums []int, lower int, upper int) [][]int {
    if len(nums) == 0 {
        return [][]int{{lower, upper}}
    }
    out := [][]int{}
    if nums[0] > lower {
        out = append(out, []int{lower, nums[0]-1})
    }
    for i := 1; i < len(nums); i++ {
        prev := nums[i-1]
        curr := nums[i]
        start, end := math.MinInt64, math.MinInt64
        for prev+1 < curr {
            if start == math.MinInt64 {start = prev+1}
            prev++
        }
        end = prev
        if start != math.MinInt64 {
            out = append(out, []int{start, end})
        }
    }
    if nums[len(nums)-1] != upper {
        out = append(out, []int{nums[len(nums)-1]+1, upper})
    }
    return out
}