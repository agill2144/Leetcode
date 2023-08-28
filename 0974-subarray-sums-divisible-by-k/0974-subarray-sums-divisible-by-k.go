func subarraysDivByK(nums []int, k int) int {
    remainderFreq := map[int]int{0:1}
    sum := 0
    count := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        remainder := sum % k
        if remainder < 0 {remainder += k}
        val, ok := remainderFreq[remainder]
        if ok {
            count += val
        }
        remainderFreq[remainder]++
    }
    return count
}