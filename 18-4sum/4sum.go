func fourSum(nums []int, target int) [][]int {
    set := map[[4]int]bool{}
    out := [][]int{}
    for i := 0; i < len(nums); i++ {
        iTarget := target-nums[i]
        for j := i+1; j < len(nums); j++ {
            jTarget := iTarget-nums[j]
            s := map[int]bool{}
            for k := j+1; k < len(nums); k++ {
                diff := jTarget-nums[k]
                if s[diff] {
                    t := []int{nums[i], nums[j], nums[k], diff}
                    sort.Ints(t)
                    tmp := [4]int{t[0],t[1],t[2],t[3]}
                    if !set[tmp] {
                        out = append(out, t)
                        set[tmp] = true
                    }
                }
                s[nums[k]] = true
            }
        }
    }
    return out
}