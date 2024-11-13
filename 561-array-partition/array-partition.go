func arrayPairSum(nums []int) int {
    sort.Ints(nums)
    totalSum := 0
    for i := 0; i < len(nums); i+=2 {
        totalSum += min(nums[i], nums[i+1])
    }
    return totalSum
}