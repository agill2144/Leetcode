func findMissingRanges(nums []int, lower int, upper int) [][]int {
    if len(nums) == 0 {return [][]int{{lower, upper}}}
    ptr := 0
    out := [][]int{}
    n := len(nums)

    if nums[ptr] != lower {
        out = append(out, []int{lower, nums[ptr]-1})
    }

    for ptr < n-1 {
        curr := nums[ptr]
        next := nums[ptr+1]
        if curr+1 != next {
            out = append(out, []int{curr+1, next-1})
        }
        ptr++
    }
    if nums[n-1] != upper {
        out = append(out, []int{nums[n-1]+1, upper})
    }

    return out
}