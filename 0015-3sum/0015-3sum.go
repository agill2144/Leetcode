func threeSum(nums []int) [][]int {
    set := map[[3]int]struct{}{}
    out := [][]int{}
    for i := 0; i < len(nums); i++ {
        target := 0-nums[i]
        numSet := map[int]struct{}{}
        for j := i+1; j < len(nums); j++ {
            comp := target-nums[j]
            if _, ok := numSet[comp]; ok {
                
                tmp := []int{nums[i], nums[j], comp}
                sort.Ints(tmp)
                if _, ok := set[[3]int{tmp[0], tmp[1], tmp[2]}]; !ok {
                    set[[3]int{tmp[0], tmp[1], tmp[2]}] = struct{}{}
                    out = append(out, tmp)
                }
            }
            numSet[nums[j]] = struct{}{}
        }
    }
    return out
}