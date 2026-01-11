func findMaxConsecutiveOnes(nums []int) int {
    out := 0
    count := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 1 {
            count++
            out = max(out, count)
        } else {
            count = 0
        }
    }
    return out
}