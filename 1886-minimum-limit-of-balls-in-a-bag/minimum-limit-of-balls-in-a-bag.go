// k = max num in nums array
// n = len of nums array
// tc = o(logk * n)
// sc = o(1)
func minimumSize(nums []int, maxOperations int) int {
    left := 1
    right := slices.Max(nums)
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        ops := 0
        for i := 0; i < len(nums); i++ {
            ops += int(math.Ceil(float64(nums[i])/float64(mid)))-1
        }
        if ops > maxOperations {
            left = mid+1
        } else {
            ans = mid
            right = mid-1
        }
    }
    return ans
}