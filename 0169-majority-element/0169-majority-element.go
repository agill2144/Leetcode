func majorityElement(nums []int) int {
    num := math.MinInt64
    count := 0
    for i := 0; i < len(nums); i++ {
        if count == 0 {
            num = nums[i]
            count = 1
        } else if nums[i] == num {
            count++
        } else {
            count--
        }
    }
    count = 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == num {count++}
        if count == (len(nums)/2)+1 {return num}
    }
    return -1
}