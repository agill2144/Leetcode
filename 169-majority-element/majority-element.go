func majorityElement(nums []int) int {
    count := 0 
    ele := math.MaxInt64
    for i := 0; i < len(nums); i++ {
        if count == 0 {
            count++
            ele = nums[i]
        } else if nums[i] == ele {
            count++
        } else {
            count--
        }
    }
    return ele
}