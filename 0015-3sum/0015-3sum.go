func threeSum(nums []int) [][]int {
    set := map[[3]int]struct{}{}
    out := [][]int{}
    for i := 0; i < len(nums); i++ {
        target := 0-nums[i]
        idxMap := map[int]int{}
        for j := i+1; j < len(nums); j++ {
            complement := target-nums[j]
            idx, ok := idxMap[complement]
            if ok {
                tmp := []int{nums[i], nums[idx], nums[j]}
                sort.Ints(tmp)
                if _, ok := set[[3]int{tmp[0], tmp[1], tmp[2]}]; !ok {
                    set[[3]int{tmp[0], tmp[1], tmp[2]}] = struct{}{}
                    out = append(out, tmp)
                }
            } else {
                idxMap[nums[j]] = j
            }
        }
    }
    return out
}