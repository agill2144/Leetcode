func subarraysWithKDistinct(nums []int, k int) int {
    if k == 0 {return 0}
    return atMostK(nums, k) - atMostK(nums, k-1)    
}

func atMostK(nums []int, k int) int {
    count := 0
    left := 0
    freq := map[int]int{}
    for i := 0; i < len(nums); i++ {
        freq[nums[i]]++
        for len(freq) > k {
            freq[nums[left]]--
            if freq[nums[left]] == 0 {delete(freq, nums[left])}
            left++
        }
        count += (i-left+1)
    }
    return count
}