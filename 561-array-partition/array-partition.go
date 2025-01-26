func arrayPairSum(nums []int) int {
    sort.Ints(nums)
    total := 0
    for i := len(nums)-2; i >= 0; i -= 2 {
        total += min(nums[i], nums[i+1])
    }
    return total
}

// 1,4,3,2
// 1,2,3,4