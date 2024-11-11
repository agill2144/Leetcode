func findPairs(nums []int, k int) int {
    set := map[[2]int]struct{}{}
    idx := map[int]int{}
    for i := 0; i < len(nums); i++ {
        c1 := nums[i]-k
        _, ok := idx[c1]
        if ok {
            tmp := [2]int{nums[i], c1}
            if c1 < nums[i] {tmp = [2]int{c1, nums[i]}}
            set[tmp] = struct{}{}
        }
        c2 := nums[i]+k
        _, ok = idx[c2]
        if ok {
            tmp := [2]int{nums[i], c2}
            if c2 < nums[i] {tmp = [2]int{c2, nums[i]}}
            set[tmp] = struct{}{}
        }
        idx[nums[i]] = i
    }
    return len(set)
}