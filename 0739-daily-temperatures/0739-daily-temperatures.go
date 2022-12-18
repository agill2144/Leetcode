func dailyTemperatures(temperatures []int) []int {
    st := [][]int{}
    out := make([]int, len(temperatures))
    
    for i := 0; i < len(temperatures); i++ {
        for len(st) != 0 && temperatures[i] > st[len(st)-1][1] {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            idx := top[0]
            out[idx] = i-idx
        }
        st = append(st, []int{i, temperatures[i]})
    }
    return out
}