func maxFrequency(nums []int, k int) int {
    sort.Ints(nums)
    n := len(nums)
    right := n-1
    sum := nums[right]
    maxFreq := 1
    for i := n-2; i >= 0; i-- {
        rightVal := nums[right]
        currVal := nums[i]
        sum += currVal
        winSize := right-i+1
        desired := rightVal * winSize
        diff := desired-sum
        if diff <= k {
            maxFreq = max(maxFreq, winSize)
        } else {
            sum -= nums[right]
            right--
        }
    }
    return maxFreq
}