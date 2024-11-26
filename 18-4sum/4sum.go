func fourSum(nums []int, target int) [][]int {
    set := map[[4]int]bool{}
    out := [][]int{}
    n := len(nums)
    // n = 6 
    // we need atleast 4 elements
    // n-4 = 6-4 = 2 is the last ith idx
    for i := 0; i < n; i++ {
        iTarget := target - nums[i]
        for j := i+1; j < n; j++ {
            jTarget := iTarget-nums[j]
            seen := map[int]bool{}
            for k := j+1; k < n; k++ {
                diff := jTarget - nums[k]
                if seen[diff] {
                    tmp := []int{nums[i], nums[j],diff , nums[k]}
                    sort.Ints(tmp)
                    tmp2 := [4]int{tmp[0],tmp[2],tmp[2],tmp[3]}
                    if !set[tmp2] {
                        set[tmp2] = true
                        out = append(out, tmp)
                    }
                }
                seen[nums[k]] = true
            }
        }
    }
    return out
}