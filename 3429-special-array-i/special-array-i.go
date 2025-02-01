func isArraySpecial(nums []int) bool {
    if len(nums) <= 1 {return true}
    for i := 1; i < len(nums); i++ {
        if nums[i] % 2 == 0 && nums[i-1] % 2 == 0 {return false}
        if nums[i] % 2 != 0 && nums[i-1] % 2 != 0 {return false}
    }
    return true
}