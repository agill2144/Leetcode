func find132pattern(nums []int) bool {
    // nums[i] < nums[k] < nums[j]
    // nums[j] is supposed to be the largest element
    // nums[k] is supposed to be the 2nd largest element , its idx > idx of j
    // nums[i] is supposed to be smallest and smaller than nums[j]. its idx < idx of j
    st := []int{} // indices, acting as nums[j]
    n := len(nums)
    numsK := math.MinInt64
    numsJ := 0
    for i := n-1; i >= 0; i-- {
        if nums[i] < numsK {return true}
        numsJ = nums[i]
        for len(st) != 0 && numsJ > nums[st[len(st)-1]] {
            numsK = nums[st[len(st)-1]]
            st = st[:len(st)-1]
        }
        st = append(st, i)
    }
    return false
}