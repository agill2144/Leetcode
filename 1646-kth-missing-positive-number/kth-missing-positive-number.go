func findKthPositive(nums []int, k int) int {
    n := len(nums)
    left := 0
    right := n-1
    for left <= right {
        mid := left + (right-left)/2
        cIdx := nums[mid]-1
        missOnLeft := cIdx - mid
        if missOnLeft < k {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    if right < 0 {return k}
    missOnLeft := (nums[right]-1) - right
    k -= missOnLeft
    return nums[right] + k
}