func findMaxConsecutiveOnes(nums []int) int {
    if nums == nil || len(nums) == 0 {
        return 0
    }
    max := 0
    count := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 1{
            count++
        } else {
            max = int(math.Max(float64(count), float64(max)))
            count = 0
        }
    }
    return int(math.Max(float64(count), float64(max)))
}