func findMissingRanges(nums []int, lower int, upper int) [][]int {
    if len(nums) == 0 {
        return [][]int{{lower, upper}}
    }
    out := [][]int{}
    if lower < nums[0] {
        out = append(out, []int{lower,nums[0]-1})
    }
    for i := 1; i < len(nums); i++ {
        prev := nums[i-1]
        curr := nums[i]
        if prev + 1 == curr {continue}
        
        start := prev+1
        end := curr-1
        if start >= lower && end <= upper {out = append(out, []int{start, end})}
    }
    if nums[len(nums)-1] < upper {
        out = append(out, []int{nums[len(nums)-1]+1, upper})
    }
    return out
}