func sortColors(nums []int)  {
    z,o,t := 0,0,0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {z++}
        if nums[i] == 1 {o++}
        if nums[i] == 2 {t++}
    }
    for i := 0; i < len(nums); i++ {
        if z > 0 {
            nums[i] = 0; z--
        } else if o > 0 {
            nums[i] = 1; o--
        } else {
            nums[i] = 2; t--
        }
    }
}