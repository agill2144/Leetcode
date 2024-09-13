func subarraysWithKDistinct(nums []int, k int) int {
    if k <= 0 {return 0}
    return atMostK(nums, k) - atMostK(nums, k-1)
}

func atMostK(nums []int, k int) int {
    n := len(nums)
    freq := map[int]int{}
    left := 0
    count := 0
    for i := 0; i < n; i++ {
        freq[nums[i]]++
        for len(freq) > k && left <= i {
            freq[nums[left]]--
            if freq[nums[left]] == 0 {delete(freq, nums[left])}
            left++
        }
        count += (i-left+1)
    }
    return count
}