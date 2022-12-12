func arrayPairSum(nums []int) int {
    sort.Ints(nums)
    sum := 0
    for i := 0; i < len(nums); i+=2 {
        sum += int(math.Min(float64(nums[i]), float64(nums[i+1])))
    }
    return sum
}