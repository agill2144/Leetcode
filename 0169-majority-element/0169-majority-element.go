func majorityElement(nums []int) int {
    count := 0
    ele := math.MinInt64
    for i := 0; i < len(nums); i++ {
        if count == 0 {
            count = 1
            ele = nums[i]
        } else if nums[i] == ele {
            count++
        } else {
            count--
        }
    }
    return ele
}