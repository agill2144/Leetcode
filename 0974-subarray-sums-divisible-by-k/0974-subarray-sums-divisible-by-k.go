func subarraysDivByK(nums []int, k int) int {
    remCount := map[int]int{0:1}
    sum := 0
    count := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        rem := sum % k
        if rem < 0 {rem += k}
        c, ok := remCount[rem]
        if ok {
            count+=c
        }
        remCount[rem]++
    }
    return count
    
}