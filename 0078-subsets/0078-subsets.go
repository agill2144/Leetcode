func subsets(nums []int) [][]int {
    out := [][]int{{}}
    for i := 0; i < len(nums); i++ {
        num := nums[i]
        for _, nestedList := range out {
            newL := make([]int, len(nestedList))
            copy(newL, nestedList)
            newL = append(newL, num)
            out = append(out, newL)
        }
    }
    return out
}