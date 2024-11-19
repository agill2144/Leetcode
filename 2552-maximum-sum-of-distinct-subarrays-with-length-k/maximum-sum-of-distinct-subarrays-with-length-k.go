func maximumSubarraySum(nums []int, k int) int64 {
    freq := map[int]int{}
    var maxSum int64 = 0
    var sum int64 = 0
    left := 0
    for i := 0; i < len(nums); i++ {
        sum += int64(nums[i])
        freq[nums[i]]++
        if i-left+1 == k {
            if len(freq) == k {
                maxSum = max(maxSum, sum)
                sum -= int64(nums[left])
                freq[nums[left]]--
                if freq[nums[left]] == 0 {delete(freq, nums[left])}
                left++
            } else {
                sum -= int64(nums[left])
                freq[nums[left]]--
                if freq[nums[left]] == 0 {delete(freq, nums[left])}
                left++
            }   
        }
        
    }
    return maxSum
}