func sortColors(nums []int)  {
    n := len(nums)
    zero := 0
    two := n-1
    i := 0
    for i <= two {
        if nums[i] == 2 {
            nums[two], nums[i] = nums[i], nums[two]
            two--
        } else if nums[i] == 0 {
            nums[zero], nums[i] = nums[i], nums[zero]
            zero++
            i++
        } else {
            i++
        }
    }
    
}