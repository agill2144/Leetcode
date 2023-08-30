func subarraysDivByK(nums []int, k int) int {
    remaindersIdx := map[int]int{0:1}
    sum := 0
    count := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        rem := (sum % k + k) % k
        val, ok := remaindersIdx[rem]
        if ok {
            count+=val
        }
        remaindersIdx[rem]++
    }
    return count
    
}