func findDisappearedNumbers(nums []int) []int {
    set := map[int]bool{}
    for i := 0; i < len(nums); i++ {set[nums[i]] = true}
    exp := 1
    out := []int{}
    for i := exp; i <= len(nums); i++ {
        if !set[i] {out = append(out, i)}
    }
    return out
}