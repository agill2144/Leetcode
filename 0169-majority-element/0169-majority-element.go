func majorityElement(nums []int) int {
    count := 0
    num := math.MinInt64
    for i := 0; i < len(nums); i++ {
        if count == 0 {
            count = 1
            num = nums[i]
        } else if nums[i] == num {
            count++
        } else {
            count--
        }
    }
    return num
}