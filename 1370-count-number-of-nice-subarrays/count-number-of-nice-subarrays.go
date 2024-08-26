func numberOfSubarrays(nums []int, k int) int {
    oddCounts := map[int]int{0:1}
    rCount := 0
    total := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] % 2 != 0 {
            rCount++
        }
        diff := rCount - k
        total += oddCounts[diff]
        oddCounts[rCount]++
    }
    return total
}