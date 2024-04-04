// nge pattern ( next greater on right )
func dailyTemperatures(temperatures []int) []int {
    st := []int{}
    n := len(temperatures)
    out := make([]int, n)
    for i := n-1; i >= 0; i-- {
        curr := temperatures[i]
        for len(st) != 0 && curr >= temperatures[st[len(st)-1]] {
            st = st[:len(st)-1]
        }
        if len(st) > 0 {
            topIdx := st[len(st)-1]
            out[i] = topIdx-i
        }
        st = append(st, i)
    }
    return out
}