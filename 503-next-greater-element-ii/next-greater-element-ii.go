func nextGreaterElements(nums []int) []int {
    n := len(nums)
    out := make([]int, n)
    for i := 0; i < n; i++ {out[i] = -1}

    // [1,  2, 1]
    // [-1,-1,-1]

    st := [][]int{} // [ [val, idx] ]
    for i := 0; i < 2*n; i++ {
        curr := nums[i%n]
        for len(st) != 0 && curr > st[len(st)-1][0] {
            top := st[len(st)-1]
            topIdx := top[1]
            st = st[:len(st)-1]
            out[topIdx] = curr
        }
        if i < n {
            st = append(st, []int{curr, i})
        }
    }
    return out
}