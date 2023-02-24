func subsets(nums []int) [][]int {
    out := [][]int{{}}
    for i := 0; i < len(nums); i++ {
        for _, list := range out {
            newL := make([]int, len(list))
            copy(newL, list)
            newL = append(newL, nums[i])
            out = append(out, newL)
        }
    }
    return out
}