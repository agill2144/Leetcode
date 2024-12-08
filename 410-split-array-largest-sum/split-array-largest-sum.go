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
        atMax := mid
        count := 1
        rSum := 0
        for i := 0; i < len(nums); i++ {
            rSum += nums[i]
            if rSum > atMax {
                count++
                rSum = nums[i]
            }
        }
        if count <= k {
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}