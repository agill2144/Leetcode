func numberOfSubarrays(nums []int, k int) int {
    freq := map[int]int{0:1}
    oddCount := 0
    total := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] % 2 != 0 {oddCount++}
        total += freq[oddCount-k]
        freq[oddCount]++
    }
    return total

}