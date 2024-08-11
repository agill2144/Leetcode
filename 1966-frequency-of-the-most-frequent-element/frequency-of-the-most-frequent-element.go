func maxFrequency(nums []int, k int) int {
    sort.Ints(nums)
    n := len(nums)
    right := n-1
    rSum := nums[n-1]
    maxFreq := 1
    for i := n-2; i >= 0; i-- {
        curr := nums[i]
        matchTo := nums[right]
        size := right-i+1
        totalSum := matchTo * size
        rSum += curr
        diff := totalSum - rSum 
        if diff <= k {
            maxFreq = max(maxFreq, size)
            continue
        }
        rSum -= nums[right]
        right--
    }
    return maxFreq
}