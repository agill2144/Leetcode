func findLonely(nums []int) []int {
    sort.Ints(nums)
    out := []int{}
    for i := 0; i < len(nums); i++ {
        prev := math.MinInt64
        next := math.MaxInt64
        if i-1 >= 0 {prev = nums[i-1]}
        if i+1 < len(nums) {next = nums[i+1]}
        if nums[i] + 1 != next &&
            nums[i] - 1 != prev && 
            nums[i] != prev &&
            nums[i] != next {out = append(out, nums[i])}
    }
    return out
}