// count this shit
func sortColors(nums []int)  {
    zeros, ones, twos := 0,0,0
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
        } else if twos > 0 {
            nums[i] = 2
            twos--
        }
    }
}