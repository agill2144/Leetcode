func majorityElement(nums []int) int {
    ele := math.MinInt64
    count := 0 
    n := len(nums)
    for i := 0; i < n; i++ {
        if count == 0{
            ele = nums[i]
            count = 1
        } else if ele == nums[i] {
            count++
        } else {
            count--
        }
    }
    count = 0
    for i := 0; i < n; i++ {
        if nums[i] == ele {count++}
        if count > n/2 {return ele}
    }
    panic("wtf")
}