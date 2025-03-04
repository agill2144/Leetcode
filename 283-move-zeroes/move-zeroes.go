func moveZeroes(nums []int)  {
    nonZero := 0
    n := len(nums)
    for i := 0; i < n; i++ {
        if nums[i] != 0 {
            nums[i], nums[nonZero] = nums[nonZero], nums[i]
            nonZero++
        }
    }
}