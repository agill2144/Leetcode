func majorityElement(nums []int) int {
    ele := math.MinInt64
    count := 0
    for i := 0; i < len(nums); i++ {
        if count == 0 {
            count = 1
            ele = nums[i]
        } else if ele == nums[i] {
            count++
        } else {
            count--
        }
    }
    return ele
}