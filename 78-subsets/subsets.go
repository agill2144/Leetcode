func subsets(nums []int) [][]int {
    out := [][]int{{}}
    for i := 0; i < len(nums); i++ {
        for _, ele := range out {
            newP := make([]int, len(ele))
            copy(newP, ele)
            newP = append(newP, nums[i])
            out = append(out, newP)
        }
    }
    return out
}