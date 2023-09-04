
func subarraysDivByK(nums []int, k int) int {
    remainders := map[int]int{0:1}
    sum := 0
    count := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        r := sum % k
        if r < 0 {r += k}
        val := remainders[r]
        count += val
        remainders[r]++
    }
    return count 
}