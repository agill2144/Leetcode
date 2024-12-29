func subarraysDivByK(nums []int, k int) int {
    rem := map[int]int{0:1}
    count := 0
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        r := sum % k
        if r < 0 {r += k}
        count += rem[r]
        rem[r]++
    }
    return count
}