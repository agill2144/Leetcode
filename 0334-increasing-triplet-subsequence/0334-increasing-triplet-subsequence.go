func increasingTriplet(nums []int) bool {
    first := math.MaxInt64
    second := math.MaxInt64
    for i := 0; i < len(nums); i++ {
        if nums[i] <= first {
            first = nums[i]
        } else if nums[i] <= second {
            second = nums[i]
        } else if first < second && second < nums[i] {
            return true
        }
    }
    return false
}