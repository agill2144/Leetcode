func findMaxConsecutiveOnes(nums []int) int {
    maxCount := 0
    ones := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 1 {
            ones++
            maxCount = max(maxCount, ones)
            continue
        }
        ones = 0
    }
    return maxCount
}