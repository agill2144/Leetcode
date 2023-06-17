func findMissingRanges(nums []int, lower int, upper int) [][]int {
    if len(nums) == 0 {
        return [][]int{{lower, upper}}
    }
    out := [][]int{}
    if nums[0] > lower {
        out = append(out, []int{lower, nums[0]-1})
    }
    
    for i := 1; i < len(nums); i++ {
        curr := nums[i]
        actualPrev := nums[i-1]
        expectedPrev := curr-1
        if expectedPrev != actualPrev {
            out = append(out, []int{actualPrev+1,expectedPrev})
        }
    }
    if nums[len(nums)-1] != upper {
        out = append(out, []int{nums[len(nums)-1]+1, upper})
    }

    return out
}