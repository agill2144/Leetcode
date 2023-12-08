func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    set := map[[4]int]struct{}{}
    out := [][]int{}
    for i := 0; i < len(nums); i++ {
        if i != 0 && nums[i] == nums[i-1] {continue}
        t1 := target-nums[i]
        for j := i+1; j < len(nums); j++ {
            t2 := t1 - nums[j]
            left := j+1
            right := len(nums)-1
            for left < right {
                total := nums[left] + nums[right]
                if total == t2 {
                    tmp := [4]int{nums[i], nums[j], nums[left], nums[right]}
                    if _, ok := set[tmp]; !ok {
                        set[tmp] = struct{}{}
                        out = append(out, tmp[:])
                    }
                    left++
                    for left < right && nums[left] == nums[left-1] {left++}
                    right--
                    for right > left && nums[right] == nums[right+1] {right--}
                } else if total > t2 {
                    right--
                } else {
                    left++
                }
            }
        }
    }
    return out
}