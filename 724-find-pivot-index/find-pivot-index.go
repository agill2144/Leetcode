func pivotIndex(nums []int) int {
    prefixSum := []int{}
    rSum := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        prefixSum = append(prefixSum, rSum)
    }
    for i := 0; i < len(nums); i++ {
        total := prefixSum[len(prefixSum)-1]
        left := 0
        if i-1 >= 0 {left = prefixSum[i-1]}
        right :=  total-prefixSum[i]
        if left == right {return i}     
    }
    return -1
}