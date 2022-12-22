func majorityElement(nums []int) int {
    if nums == nil || len(nums) == 0 {return 0}
    if len(nums) == 1{ return 1}

    sort.Ints(nums)
    num := nums[0]
    count := 1
    
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
            count++
        } else {
            if count > len(nums)/2 {
                return num
            }
            num = nums[i]
            count = 1
        }
    }
    return num
}