func findMaxK(nums []int) int {
    ans := -1
    set := map[int]struct{}{}
    for i := 0; i < len(nums); i++ {
        isPositive := nums[i] >= 0
        if isPositive {
            if _, ok := set[nums[i]*-1]; ok {
                ans = max(ans, nums[i])
            }
        } else {
            if _, ok := set[nums[i]*-1]; ok {
                ans = max(ans, nums[i]*-1)
            }
        }
        set[nums[i]] = struct{}{}
    }
    return ans
}