func containsDuplicate(nums []int) bool {
    set := map[int]struct{}{}
    for _, ele := range nums {
        if _, ok := set[ele]; ok {
            return true
        }
        set[ele] = struct{}{}
    }
    return false
}