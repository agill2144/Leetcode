func countSubarrays(nums []int, k int64) int64 {
    count :=  0
    var sum int64 = 0
    left := 0
    for i := 0; i < len(nums); i++ {
        sum += int64(nums[i])
        size := i-left+1
        for sum * int64(size) >= k {
            sum -= int64(nums[left])
            left++
            size = i-left+1
        }
        count += (i-left+1)
    }
    return int64(count)
}