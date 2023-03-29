func maxSubArray(nums []int) int {
    if nums == nil || len(nums) == 0 {return 0}
    
    curr := nums[0]
    maxTotal := nums[0]
    
    for i := 1; i < len(nums); i++ {
        newCurr := max(nums[i], curr + nums[i])
        curr = newCurr
        if curr > maxTotal {maxTotal = curr}
    }
    return maxTotal    
}

func max(x, y int) int {
    if x > y {return x} 
    return y
}