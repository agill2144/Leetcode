// mores voting algo
func majorityElement(nums []int) int {
    count := 0
    candidate := math.MinInt64
    for i := 0; i < len(nums); i++ {
        if count == 0 {
            candidate = nums[i]
            count++
        } else if nums[i] == candidate {
            count++
        } else {
            count--
        }
    }
    return candidate
}