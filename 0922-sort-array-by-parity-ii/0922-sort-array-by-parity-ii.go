func sortArrayByParityII(nums []int) []int {
    even := 0
    i := 1
    for i < len(nums) {
        if nums[i] % 2 == 0 {
            nums[i], nums[even] = nums[even], nums[i]
            even+=2
        } else {
            i += 2
        }
    }
    return nums
}