// this is not a medium ... 
// this is borderline hard .. wtf
// https://www.youtube.com/watch?v=iOqH_JnXIOQ
// time = o(nlogn) + o(n * logn)
func maxFrequency(nums []int, k int) int {
    sort.Ints(nums)
    prefixSum := make([]int, len(nums))
    prefixSum[0] = nums[0]
    for i := 1; i < len(nums); i++ {
        prefixSum[i] = prefixSum[i-1]+nums[i]
    }
    maxFreq := 0
    for i := 0; i < len(nums); i++ {
        left := 0
        right := i
        target := nums[i]
        for left <= right {
            mid := left + (right-left)/2
            windowSize := i-mid+1
            desiredSum := windowSize * target
            leftSumToRemove := 0
            if mid-1 >= 0 {leftSumToRemove = prefixSum[mid-1]}
            actualSum := prefixSum[i]-leftSumToRemove
            ops := desiredSum - actualSum
            if ops <= k {
                maxFreq = max(maxFreq, windowSize)
                right = mid-1
            } else {
                left = mid+1
            }
        }
    }
    return maxFreq
}