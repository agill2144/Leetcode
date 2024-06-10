func find132pattern(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
    st := []int{}
    nums3 := math.MinInt64 // k
    for i := n-1; i >= 0; i-- {
        if nums[i] < nums3 {
            return true
        }

        for len(st) != 0 && nums[i] > st[len(st)-1] {
            nums3 = st[len(st)-1]
            st = st[:len(st)-1]
        }
        st = append(st, nums[i])
    }
    return false
}