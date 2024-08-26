func numberOfSubarrays(nums []int, k int) int {
    return atMostK(nums, k) - atMostK(nums, k-1)    
}

func atMostK(nums []int, k int) int {
    count := 0
    rCount := 0
    left := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] % 2 != 0 {rCount++}
        for rCount > k {
            if nums[left] % 2 != 0 {rCount--}
            left++
        }
        count += (i-left+1)
    }
    return count
}