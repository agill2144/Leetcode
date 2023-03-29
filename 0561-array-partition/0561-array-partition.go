func arrayPairSum(nums []int) int {
    sort.Ints(nums)
    total := 0
    for i := 0; i < len(nums); i+=2 {
        total += min(nums[i], nums[i+1])
    }
    return total
}

func min(x, y int) int {
    if x < y {return x}
    return y
}