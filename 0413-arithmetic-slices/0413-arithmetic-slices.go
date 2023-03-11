func numberOfArithmeticSlices(nums []int) int {
    count := 0
    for i := 0; i < len(nums)-1; i++ {
        diff := nums[i+1]-nums[i]
        for j := i+2; j < len(nums); j++ {
            if nums[j]-nums[j-1] == diff {
                if j-i+1 >= 3{
                    count++; continue
                }
            }
            break
        }
    }
    return count
}