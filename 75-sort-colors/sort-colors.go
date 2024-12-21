func sortColors(nums []int)  {
    zero := 0
    two := len(nums)-1
    i := 0
    for i <= two {
        if nums[i] == 2 {
            nums[i], nums[two] = nums[two], nums[i]
            two--
        } else if nums[i] == 0 {
            nums[i], nums[zero] = nums[zero], nums[i]
            zero++
            i++
        } else {
            i++
        }
    }

}