/*
    largest sum of any subarray is minimized.
    - binary search hint
    - mid = atMax
    - use "minimize" to search on correct side of mid if mid is an answer
    - mid is answer when it took <= k-1 split to satisfy atMax sum per chunk
*/
func splitArray(nums []int, k int) int {
    left := -1
    right := 0
    for i := 0; i < len(nums); i++ {
        left = max(left, nums[i])
        right += nums[i]
    }
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        rSum := 0
        splits := 0
        for i := 0; i < len(nums); i++ {
            rSum += nums[i]
            if rSum > mid {
                rSum = nums[i]
                splits++
            }
        }
        if splits <= k-1 {
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}