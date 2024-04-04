func dailyTemperatures(temperatures []int) []int {
    st := [][]int{} // [temp, idx]
    n := len(temperatures)
    out := make([]int, n)
    for i := 0; i < n; i++ {
        curr := temperatures[i]
        for len(st) != 0 && curr > st[len(st)-1][0] {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            idx := top[1]
            out[idx] = i-idx
        }
        st = append(st, []int{curr, i})
    }
    return out
}