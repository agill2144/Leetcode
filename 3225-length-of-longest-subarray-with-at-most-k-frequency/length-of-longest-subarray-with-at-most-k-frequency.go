func maxSubarrayLength(nums []int, k int) int {
    if k <= 0 {return 0}
    if k >= len(nums) {return len(nums)}
    left := 0
    freq := map[int]int{}
    size := 0
    for i := 0; i < len(nums); i++ {
        freq[nums[i]]++
        for freq[nums[i]] > k {
            freq[nums[left]]--
            if freq[nums[left]] == 0 {delete(freq, nums[left])}
            left++
        }
        size = max(size, i-left+1)
    }
    return size
}