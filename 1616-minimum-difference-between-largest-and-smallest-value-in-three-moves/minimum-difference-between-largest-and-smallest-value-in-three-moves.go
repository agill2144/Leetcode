func minDifference(nums []int) int {
    if len(nums) <= 4 {return 0}
    sort.Ints(nums)
    minDiff := math.MaxInt64
    // remove 3 from left hand side
    minDiff = min(minDiff, nums[len(nums)-1]-nums[3])
    // remove 3 from right hand side
    minDiff = min(minDiff, nums[len(nums)-3-1]-nums[0])
    // remove 1 from left and 2 from right
    minDiff = min(minDiff, nums[len(nums)-2-1]-nums[1])
    // remove 2 from left and 1 from right
    minDiff = min(minDiff, nums[len(nums)-1-1]-nums[2])
    return minDiff
}