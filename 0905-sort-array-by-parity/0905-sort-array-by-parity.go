func sortArrayByParity(nums []int) []int {
    evenCollector := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] % 2 == 0 {
            nums[i], nums[evenCollector] = nums[evenCollector], nums[i]
            evenCollector++
        }
    }
    return nums
}