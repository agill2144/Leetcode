func numberOfSubarrays(nums []int, k int) int {
    if k < 1 {return 0}
    return atMostK(nums, k) - atMostK(nums, k-1)
}

func atMostK(nums []int, k int) int {
    n := len(nums)
    total := 0
    left := 0
    rCount := 0
    for i := 0; i < n; i++ {
        if nums[i] % 2 != 0 {rCount++}
        for rCount > k {
            if nums[left] % 2 != 0 {rCount--}
            left++
        }
        total += (i-left+1)
    }
    return total
}

// running sum pattern:
// time = o(n)
// space = o(n)
// func numberOfSubarrays(nums []int, k int) int {
//     freq := map[int]int{0:1}
//     rCount := 0
//     total := 0
//     for i := 0; i < len(nums); i++ {
//         if nums[i] % 2 != 0 {rCount++}
//         total += freq[rCount-k]
//         freq[rCount]++
//     }
//     return total
// }