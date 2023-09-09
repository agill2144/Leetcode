func fourSum(nums []int, target int) [][]int {
    set := map[[4]int]struct{}{}
    out := [][]int{}
    for i := 0; i < len(nums); i++ {
        iTarget := target-nums[i]
        
        for j := i+1; j < len(nums); j++ {
            jTarget := iTarget-nums[j]

            s := map[int]struct{}{}
            for k := j+1; k < len(nums); k++ {
                diff := jTarget - nums[k]
                if _, ok := s[diff]; ok {
                    tmp := []int{nums[i],nums[j],nums[k], diff}
                    sort.Ints(tmp)
                    newTmp := [4]int{tmp[0],tmp[1], tmp[2],tmp[3]}
                    if _, ok := set[newTmp]; !ok {
                        set[newTmp] = struct{}{}
                        out = append(out, tmp)
                    }                    
                }
                s[nums[k]] = struct{}{}
            }
        }
    }
    return out
}