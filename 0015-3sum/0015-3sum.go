func threeSum(nums []int) [][]int {
    set := map[[3]int]struct{}{}
    out := [][]int{}
    for i := 0; i < len(nums); i++ {
        target := 0-nums[i]
        comps := map[int]struct{}{}
        for j := i+1; j < len(nums); j++ {
            complement := target-nums[j]
            if _, ok := comps[complement]; ok {
                tmp := []int{nums[i], nums[j], complement}
                sort.Ints(tmp)
                if _, seen := set[[3]int{tmp[0], tmp[1], tmp[2]}]; !seen {
                    set[[3]int{tmp[0], tmp[1], tmp[2]}] = struct{}{}
                    out = append(out, tmp)
                }
            }
            comps[nums[j]] = struct{}{}
         }
    }
    return out
}