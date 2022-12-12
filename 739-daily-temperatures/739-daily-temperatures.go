func dailyTemperatures(temperatures []int) []int {
    out := make([]int, len(temperatures))
    st := []int{}
    for i := 0; i < len(temperatures); i++ {
        for len(st) != 0 && temperatures[i] > temperatures[st[len(st)-1]] {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            out[top] = i - top
        }
        st = append(st, i)
    }
    return out
}