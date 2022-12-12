func sortColors(nums []int)  {
    tPtr := len(nums)-1
    zPtr := 0
    i := 0
    for i <= tPtr {
        if nums[i] == 0 {
            nums[i], nums[zPtr] = nums[zPtr], nums[i]
            i++; zPtr++
        } else if nums[i] == 2 {
            nums[i], nums[tPtr] = nums[tPtr], nums[i]
            tPtr--
        } else {
            i++
        }
    }
}