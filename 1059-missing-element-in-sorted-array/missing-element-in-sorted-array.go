func missingElement(nums []int, k int) int {
    offSet := nums[0]
    n := len(nums)
    left := 0
    right := n-1
    for left <= right {
        mid := left + (right-left)/2
        correctVal := mid+offSet
        missing := nums[mid]-correctVal
        if missing >= k {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    missing := nums[right]-(right+offSet)
    return nums[right] + (k-missing)
}