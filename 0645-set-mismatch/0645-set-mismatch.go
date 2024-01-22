func findErrorNums(nums []int) []int {
    missing := 1
    duplicate := -1
    for i := 0; i < len(nums); i++ {
        idx := abs(nums[i])-1
        if nums[idx] < 0 {
            duplicate = abs(nums[i])
        } else {
            nums[idx] *= -1
        }
    }
    for i := 0; i < len(nums); i++ {
        if nums[i] > 0 {
            missing = i+1
            break
        }
    }
    return []int{duplicate, missing}
}

func abs(x int) int {
    if x < 0 {return x*-1}
    return x
}