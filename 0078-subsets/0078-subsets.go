func subsets(nums []int) [][]int {
    out := [][]int{{}}
    for i := 0; i < len(nums); i++ {
        for _, currList := range out {
            newL := make([]int, len(currList))
            copy(newL, currList)
            newL = append(newL, nums[i])
            out = append(out, newL)
        }
    }
    return out
}