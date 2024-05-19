func nextGreaterElements(nums []int) []int {
    n := len(nums)
    st := []int{} // indices
    out := make([]int, n)
    for i := 0; i < n; i++ {out[i] = -1}

    for i := 0; i < 2 * n; i++ {
        curr := nums[i%n]
        for len(st) != 0 && curr > nums[st[len(st)-1]] {
            // process top
            top := st[len(st)-1]
            st = st[:len(st)-1]
            out[top] = curr
        }
        st = append(st,i%n)
    }
    return out
}