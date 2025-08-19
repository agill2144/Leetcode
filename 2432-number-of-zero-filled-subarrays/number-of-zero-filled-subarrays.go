func zeroFilledSubarray(nums []int) int64 {
    var count int64
    left := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            left = i+1
        } else {
            count += int64(i-left+1)
        }
    }
    return count
}