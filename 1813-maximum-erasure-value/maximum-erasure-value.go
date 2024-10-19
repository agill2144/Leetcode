func maximumUniqueSubarray(nums []int) int {
    freq := map[int]int{}
    maxScore := 0
    rSum := 0
    left := 0
    for i := 0; i < len(nums); i++ {
        for freq[nums[i]] > 0 {
            freq[nums[left]]--
            rSum -= nums[left]
            if freq[nums[left]] == 0 {delete(freq, nums[left])}
            left++            
        }
        freq[nums[i]]++
        rSum += nums[i]
        maxScore = max(maxScore, rSum)
    }
    return maxScore
}