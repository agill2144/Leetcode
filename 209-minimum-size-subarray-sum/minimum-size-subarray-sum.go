func minSubArrayLen(target int, nums []int) int {
    left := 0
    sum := 0
    minSize := len(nums)+1

    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        for sum >= target {
            minSize = min(minSize, i-left+1)
            sum -= nums[left]
            left++
        }
    }

    if minSize == len(nums)+1 {return 0}
    return minSize
}