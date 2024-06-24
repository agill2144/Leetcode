func sortColors(nums []int)  {
    zeros := 0
    ones := 0
    twos := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            zeros++
        } else if nums[i] == 1 {
            ones++
        } else {
            twos++
        }
    }
    for i := 0; i < len(nums); i++ {
        if zeros > 0 {
            nums[i] = 0
            zeros--
        } else if ones > 0 {
            nums[i] = 1
            ones--
        } else {
            nums[i] = 2
            twos--
        }
    }
    
}