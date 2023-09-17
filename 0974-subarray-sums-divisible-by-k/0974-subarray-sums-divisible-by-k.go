func subarraysDivByK(nums []int, k int) int {
    count := 0
    rem := map[int]int{0:1}
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        r := sum % k
        if r < 0 {r += k}
        val, ok := rem[r]
        if ok {
            count += val
        }
        rem[r]++
    }
    return count
}