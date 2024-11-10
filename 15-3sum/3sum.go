func threeSum(nums []int) [][]int {
    set := map[[3]int]bool{}
    out := [][]int{}
    for i := 0; i < len(nums); i++ {
        target := 0-nums[i]
        comp := map[int]bool{}
        for j := i+1; j < len(nums); j++ {
            diff := target - nums[j]
            if comp[diff] {
                trips := []int{nums[i],diff, nums[j]}
                sort.Ints(trips)
                tmp := [3]int{trips[0],trips[1], trips[2]}
                if !set[tmp] {
                    out = append(out, trips)
                    set[tmp] = true
                }
            }
            comp[nums[j]] = true
        }
    }
    return out
}