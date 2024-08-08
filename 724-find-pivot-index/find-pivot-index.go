func pivotIndex(nums []int) int {
    prefixSum := make([]int, len(nums))   
    rSum := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        prefixSum[i] = rSum
    }
    for i := 0; i < len(prefixSum); i++ {
        leftSum := 0
        if i-1 >= 0 {leftSum = prefixSum[i-1]}
        rightSum := prefixSum[len(prefixSum)-1] - prefixSum[i]
        if leftSum == rightSum {return i}
    }
    return -1
}