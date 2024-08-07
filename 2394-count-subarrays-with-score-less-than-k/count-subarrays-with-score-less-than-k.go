func countSubarrays(nums []int, k int64) int64 {
    // score = sum of all elements in window * len of window
    var count int64
    left := 0
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        score := int64(sum) * int64(i-left+1)
        for score >= k {
            sum -= nums[left]
            left++
            score = int64(sum) * int64(i-left+1)
        }
        count += int64(i-left+1)
    }
    return count
}