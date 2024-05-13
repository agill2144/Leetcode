func nextGreaterElements(nums []int) []int {
    n := len(nums)
    out := make([]int, n)
    for i := 0; i < n; i++ {out[i] = -1}
    st := []int{}
    for i := 0; i < 2*n; i++ {
        curr := nums[i%n]
        for len(st) != 0 && curr > nums[st[len(st)-1]] {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            out[top] = curr
        }
        if i >= n {continue}
        st = append(st, i%n)
    }
    return out
}