func maxAbsoluteSum(nums []int) int {
    minSoFar := 0
    maxSoFar := 0
    res := 0
    for i := 0; i < len(nums); i++ {
        // like traditional kadane's algo
        // if running sum becomes negative, we reset
        // similarly, we have to have track the traditional rSum (maxSoFar)
        // and also minSoFar, we want to keep this value as small as possible
        // therefore if minSoFar becomes > 0, we reset
        // and maxSoFar is traditional rSum from kadanes algo
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