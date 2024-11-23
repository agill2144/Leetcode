func pivotIndex(nums []int) int {
    total := 0
    for i := 0; i < len(nums); i++ {
        total += nums[i]
    }
    rSum := 0
    for i := 0; i < len(nums); i++ {
        leftSum := rSum
        // we want the sum from i+1 to n-1
        // total has everything
        // first remove leftSum from total ( this excludes ith element )
        // then remove ith element too, because remember total sum is sum of all nums
        // while leftSum is sum of nums from idx 0 to i-1 ( does not include i) 
        rightSum := total - leftSum - nums[i]
        if leftSum == rightSum {return i}
        rSum += nums[i]
    }
    return -1
}