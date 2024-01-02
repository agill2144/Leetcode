func findMatrix(nums []int) [][]int {
    out := [][]int{}
    usedIdxSet := map[int]struct{}{}
    for len(usedIdxSet) != len(nums) {

        rowSet := map[int]struct{}{}
        for i := 0; i < len(nums); i++ {
            val := nums[i]
            if _, ok := usedIdxSet[i]; ok {continue}
            if _, ok := rowSet[val]; !ok {
                usedIdxSet[i] = struct{}{}
                rowSet[val] = struct{}{}
            }
        }
        ls := []int{}
        for k, _ := range rowSet {
            ls = append(ls, k)
        }
        out = append(out, ls)
        
    }
    return out
}