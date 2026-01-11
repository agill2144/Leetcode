func moveZeroes(nums []int)  {
    nz := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[nz],nums[i] = nums[i], nums[nz]
            nz++
        }
    }
}