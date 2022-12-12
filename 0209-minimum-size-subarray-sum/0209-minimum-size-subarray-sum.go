func minSubArrayLen(target int, nums []int) int {
    left := 0
    windowSum := 0
    minSize := math.MaxInt64
    for right := 0; right < len(nums); right++ {
        windowSum += nums[right]
        for windowSum >= target {
            if right-left+1 < minSize {
                minSize = right-left+1
            }

            windowSum -= nums[left]
            left++
        }
    }
    if minSize == math.MaxInt64 {return 0}
    return minSize
}