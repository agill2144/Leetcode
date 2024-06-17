func find132pattern(nums []int) bool {
    st := []int{} // indices, acting as nums[j]
    n := len(nums)
    numsK := math.MinInt64
    for i := n-1; i >= 0; i-- {
        if nums[i] < numsK {return true}
        for len(st) != 0 && nums[i] > nums[st[len(st)-1]] {
            numsK = nums[st[len(st)-1]]
            st = st[:len(st)-1]
        }
        st = append(st, i)
    }
    return false
}