func getSneakyNumbers(nums []int) []int {
    set := map[int]struct{}{}
    out := []int{}
    for i := 0; i < len(nums); i++ {
        _, ok := set[nums[i]]
        if ok {
            out = append(out, nums[i])
        } else {
            set[nums[i]] = struct{}{}
        }
    }
    return out
}