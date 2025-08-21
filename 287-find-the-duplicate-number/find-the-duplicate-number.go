func findDuplicate(nums []int) int {
    n := len(nums)
    for i := 0; i < n; i++ {
        idx := abs(nums[i])-1
        if nums[idx] < 0 {return abs(nums[i])}
        nums[idx] *= -1        
    }
    return 0
}

func abs(x int) int {
    if x < 0 {return x * -1 }
    return x
}