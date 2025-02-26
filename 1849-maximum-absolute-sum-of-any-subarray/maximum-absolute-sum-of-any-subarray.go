func maxAbsoluteSum(nums []int) int {
    minSoFar := 0
    maxSoFar := 0
    res := 0
    for i := 0; i < len(nums); i++ {
        minSoFar = min(0, minSoFar+nums[i])
        maxSoFar = max(0, maxSoFar+nums[i])
        res = max(maxSoFar, max(abs(minSoFar),res))
    }
    return res
}

func abs(x int) int {
    if x < 0 {return x * -1}
    return x
}