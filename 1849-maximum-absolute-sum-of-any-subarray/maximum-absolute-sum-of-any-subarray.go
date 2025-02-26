func maxAbsoluteSum(nums []int) int {
    minSoFar := 0
    maxSoFar := 0
    res := 0
    for i := 0; i < len(nums); i++ {
        minSoFar += nums[i]
        if minSoFar > 0 {minSoFar = 0}
        maxSoFar += nums[i]
        if maxSoFar < 0 {maxSoFar = 0}
        res = max(maxSoFar, max(abs(minSoFar),res))
    }
    return res
}

func abs(x int) int {
    if x < 0 {return x * -1}
    return x
}