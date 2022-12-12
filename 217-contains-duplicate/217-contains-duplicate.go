func containsDuplicate(nums []int) bool {
    sort.Ints(nums)
    if len(nums) == 1 {
        return false
    }
    for i := 1; i<len(nums); i++ {
        if nums[i] == nums[i-1] {
            return true
        }
    }
    return false
}