func splitArray(nums []int, k int) int {
    left := math.MinInt64
    right := 0
    for i := 0; i < len(nums); i++ {
        right += nums[i]
        left = max(left, nums[i])
    }
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        atMax := mid
        count := 1
        sum := 0
        for j := 0; j < len(nums); j++{
            sum += nums[j]
            if sum > atMax {
                count++
                sum = nums[j]
            }
        }
        if count > k {
            left = mid+1
        } else {
            ans = mid
            right = mid-1
        }
    }
    return ans
}