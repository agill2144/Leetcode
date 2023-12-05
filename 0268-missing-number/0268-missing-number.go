func missingNumber(nums []int) int {
    set := map[int]struct{}{}
    n := len(nums)
    for i := 0; i < len(nums); i++ {
        set[nums[i]] = struct{}{}
    }
    for i := 0; i <= n; i++ {
        if _, ok := set[i]; !ok {return i}
    }
    return n
}