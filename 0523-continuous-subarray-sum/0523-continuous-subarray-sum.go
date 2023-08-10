// only do this after https://leetcode.com/problems/subarray-sum-equals-k/
func checkSubarraySum(nums []int, k int) bool {
    rSum := 0
    remainderIdx := map[int]int{0:-1}
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        remainder := rSum % k
        idx, ok := remainderIdx[remainder]
        if ok {
            if i-(idx+1)+1 >=2 {return true}
        } else {
            remainderIdx[remainder] = i
        }
    }
    return false
}