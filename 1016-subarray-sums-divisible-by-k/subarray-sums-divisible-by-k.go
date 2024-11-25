func subarraysDivByK(nums []int, k int) int {
    rems := map[int]int{0:1}
    count := 0
    rSum := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        rem := rSum % k
        if rem < 0 {rem += k}
        count += rems[rem]
        rems[rem]++
    }
    return count
}