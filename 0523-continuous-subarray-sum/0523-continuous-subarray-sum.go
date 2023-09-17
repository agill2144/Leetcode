func checkSubarraySum(nums []int, k int) bool {
    remainders := map[int]int{}
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        if sum % k == 0 && i >= 1 {return true}
        rem := sum % k 
        idx, ok := remainders[rem]
        if ok {
            if i-(idx+1)+1 >= 2 {return true}
        } else {
            remainders[rem] = i
        }
    }
    return false
}