func check(nums []int) bool {
    if len(nums) <= 1 {return true}

    startIdx := -1
    for i := 0; i < len(nums)-1; i++ {
        if nums[i] > nums[i+1] {
            startIdx = i+1
            break
        }
    }
    if startIdx == -1 {return true}
    count := 0
    n := len(nums)
    for count != n-1 {
        if nums[startIdx%n] <= nums[(startIdx+1)%n] {
            count++
            startIdx++
        } else {
            return false
        }
    }
    return true
}