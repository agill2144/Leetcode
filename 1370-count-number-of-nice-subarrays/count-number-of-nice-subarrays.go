func numberOfSubarrays(nums []int, k int) int {
    if k == 0 {return 0}
    return atMostK(nums, k) - atMostK(nums, k-1)
}

func atMostK(nums []int, k int) int {
    count := 0
    oddCount := 0
    left := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] % 2 != 0 {oddCount++}
        for left <= i && oddCount > k {
            if nums[left] % 2 != 0 {oddCount--}
            left++
        }
        count += (i-left+1)
    }
    return count
}
// func numberOfSubarrays(nums []int, k int) int {
//     freq := map[int]int{0:1}
//     oddCount := 0
//     total := 0
//     for i := 0; i < len(nums); i++ {
//         if nums[i] % 2 != 0 {oddCount++}
//         total += freq[oddCount-k]
//         freq[oddCount]++
//     }
//     return total

// }