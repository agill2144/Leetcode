func subarraysWithKDistinct(nums []int, k int) int {
    if k == 0 {return 0}
    return countLessThanOrEqualTo(nums, k) - countLessThanOrEqualTo(nums, k-1)
}

func countLessThanOrEqualTo(nums []int, k int) int {
    freq := map[int]int{}
    left := 0
    count := 0
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