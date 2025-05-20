func majorityElement(nums []int) int {
    n := len(nums)
    ele := math.MinInt64
    count := 0
    for i := 0; i < n; i++ {
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