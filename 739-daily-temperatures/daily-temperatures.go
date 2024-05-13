func dailyTemperatures(temperatures []int) []int {
    n := len(temperatures)
    out := make([]int, n)
    st := []int{} // indicies
    // ngr, process top
    for i := 0; i < n; i++ {
        curr := temperatures[i]
        for len(st) != 0 && curr > temperatures[st[len(st)-1]] {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            out[top] = i-top
        }
        st = append(st, i)
    }
    return out
}