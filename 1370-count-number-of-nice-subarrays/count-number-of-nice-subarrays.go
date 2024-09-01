func numberOfSubarrays(nums []int, k int) int {
    total := 0
    rCount := 0
    counts := map[int]int{0:1}
    for i := 0; i < len(nums); i++ {
        if nums[i] % 2 != 0 {
            rCount++
        }
        diff := rCount-k
        total += counts[diff]
        counts[rCount]++
    }
    return total
}