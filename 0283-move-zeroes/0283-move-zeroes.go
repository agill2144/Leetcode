func moveZeroes(nums []int)  {
    slow := 0
    fast := 0
    for fast < len(nums) {
        if nums[fast] != 0 {
            nums[slow], nums[fast] = nums[fast], nums[slow]
            slow++
        }
        fast++
    }
}