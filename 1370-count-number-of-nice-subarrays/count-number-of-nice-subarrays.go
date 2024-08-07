func numberOfSubarrays(nums []int, k int) int {
    oddCount := 0
    oddCounts := map[int]int{0:1}
    total := 0   
    for i := 0; i < len(nums); i++ {
        if nums[i] % 2 != 0 {oddCount++}
        remove := oddCount-k
        total += oddCounts[remove]
        oddCounts[oddCount]++
    }
    return total
}