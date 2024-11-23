func numberOfSubarrays(nums []int, k int) int {
    freq := map[int]int{0:1}
    rCount := 0
    total := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] % 2 != 0 {rCount++}
        total += freq[rCount-k]
        freq[rCount]++
    }
    return total
}