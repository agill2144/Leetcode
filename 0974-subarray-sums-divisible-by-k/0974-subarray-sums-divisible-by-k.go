func subarraysDivByK(nums []int, k int) int {
    count := 0
    freq := map[int]int{0:1}
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        rem := sum % k 
        if rem < 0 {rem += k}
        val, ok := freq[rem]
        if ok {
            count += val
        }
        freq[rem]++
    }
    return count
}