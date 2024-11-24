func findPeakElement(nums []int) int {
    // strictly greater than its neighbors
    // not >= , strictly > 
    for i := 0; i < len(nums); i++ {
        prev := math.MinInt64
        if i-1 >= 0 {prev = nums[i-1]}
        curr := nums[i]
        next := math.MinInt64
        if i+1 < len(nums) {next = nums[i+1]}
        if curr > prev && curr > next {return i}
    }
    return -1
}