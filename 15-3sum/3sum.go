func threeSum(nums []int) [][]int {
    target := 0
    set := map[[3]int]bool{}
    n := len(nums)
    out := [][]int{}
    for i := 0; i < n; i++ {
        iTarget := target - nums[i]
        set2 := map[int]bool{}
        for j := i+1; j < n; j++ {
            diff := iTarget - nums[j]
            if set2[diff] {
                tmp := []int{nums[i], nums[j], diff}
                sort.Ints(tmp)
                if !set[[3]int{tmp[0],tmp[1],tmp[2]}] {
                    out = append(out, tmp)
                    set[[3]int{tmp[0],tmp[1],tmp[2]}] = true
                }
            }
            set2[nums[j]] = true
        }
    }
    return out
}