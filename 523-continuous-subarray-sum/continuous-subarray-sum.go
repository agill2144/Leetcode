func checkSubarraySum(nums []int, k int) bool {
    remIdx := map[int]int{0:-1}
    rSum := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]        
        idx, ok := remIdx[rSum%k]
        if ok && i-idx >= 2 {return true}
        if !ok {
            remIdx[rSum%k] = i
        }
    }   
    return false
}