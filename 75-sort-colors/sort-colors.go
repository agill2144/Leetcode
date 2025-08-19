// 2 ptrs on each end collecting their values
func sortColors(nums []int)  {
    n := len(nums)
    z, t := 0, n-1
    i := 0
    for i <= t {
        if nums[i] == 0 {
            nums[i], nums[z] = nums[z], nums[i]
            i++
            z++
        } else if nums[i] == 2 {
            nums[i], nums[t] = nums[t], nums[i]
            t--
        } else {
            i++
        }
    }
}